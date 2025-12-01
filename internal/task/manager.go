package task

import (
	"context"
	"fmt"
	"slices"
	"sync"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/topvennie/spotify_organizer/internal/database/model"
	"github.com/topvennie/spotify_organizer/internal/database/repository"
	"go.uber.org/zap"
)

// Manager is the global single task manager instance
var Manager *manager

type job struct {
	task       model.Task
	status     Status
	interval   time.Duration
	lastStatus model.TaskResult
	lastError  error

	users []model.User // If it's not empty then an user triggered it and is waiting on it
}

// Manager can be used to schedule  recurring tasks in the background
// It keeps logs inside the database.
// However it does not automatically reshedule tasks after an application reboot
type manager struct {
	scheduler gocron.Scheduler
	repo      repository.Task

	mu   sync.Mutex
	jobs map[string]job
}

func newManager(repo repository.Repository) (*manager, error) {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		return nil, fmt.Errorf("failed to create new scheduler %w", err)
	}

	scheduler.Start()

	manager := &manager{
		scheduler: scheduler,
		repo:      *repo.NewTask(),
		jobs:      make(map[string]job),
	}

	if err := manager.repo.SetInactiveAll(context.Background()); err != nil {
		return nil, err
	}

	return manager, nil
}

// Add adds a new task to the manager.
// It immediately runs the task and then schedules it according to the interval.
// An unique uid is required.
// History logs (in the DB) for recurrent tasks are accessed by uid.
// If you change a task's uid then all it's history will be lost (but still in the DB)
func (m *manager) Add(ctx context.Context, newTask Task) error {
	zap.S().Infof("Adding task: %s | interval: %s", newTask.Name(), newTask.Interval())

	if _, ok := m.jobs[newTask.UID()]; ok {
		return fmt.Errorf("task %s already exists (uid: %s)", newTask.Name(), newTask.UID())
	}

	task, err := m.repo.GetByUID(ctx, newTask.UID())
	if err != nil {
		return err
	}
	if task != nil {
		// Pre-existing task
		// Update it
		task.Name = newTask.Name()
		task.Active = true
		if err := m.repo.Update(ctx, *task); err != nil {
			return err
		}
	} else {
		// New task
		// Let's create it
		task = &model.Task{
			UID:    newTask.UID(),
			Name:   newTask.Name(),
			Active: true,
		}
		if err := m.repo.Create(ctx, *task); err != nil {
			return err
		}
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	// Will immediately try to execute but it'll have to wait until the lock is released
	if _, err := m.scheduler.NewJob(
		gocron.DurationJob(newTask.Interval()),
		gocron.NewTask(m.wrap(newTask)),
		gocron.WithName(task.UID),
		gocron.WithContext(newTask.Ctx()),
		gocron.WithTags(task.UID),
		gocron.WithStartAt(gocron.WithStartImmediately()),
	); err != nil {
		return fmt.Errorf("failed to add task %+v | %w", *task, err)
	}

	m.jobs[task.UID] = job{
		task:       *task,
		status:     Waiting,
		interval:   newTask.Interval(),
		lastStatus: model.TaskSuccess,
		lastError:  nil,
		users:      []model.User{},
	}

	return nil
}

// RunByUID runs a pre existing task given a task UID.
func (m *manager) RunByUID(taskUID string, user model.User) error {
	var job gocron.Job
	for _, j := range m.scheduler.Jobs() {
		if taskUID == j.Tags()[0] {
			job = j
			break
		}
	}
	if job == nil {
		return fmt.Errorf("task with uid %s not found", taskUID)
	}

	// Set the user argument
	m.mu.Lock()
	info, ok := m.jobs[taskUID]
	if !ok {
		m.mu.Unlock()
		return fmt.Errorf("task with uid %s not found", taskUID)
	}
	info.users = append(info.users, user)
	m.jobs[taskUID] = info
	m.mu.Unlock()

	if err := job.RunNow(); err != nil {
		return fmt.Errorf("failed to run task with uid %s | %w", taskUID, err)
	}

	return nil
}

// Tasks returns all scheduled tasks
func (m *manager) Tasks() ([]Stat, error) {
	m.mu.Lock()
	jobs := m.scheduler.Jobs()
	jobsRecurring := m.jobs
	m.mu.Unlock()

	stats := make([]Stat, 0, len(jobs))

	for _, job := range jobs {
		taskUID := job.Tags()[0]

		nextRun, err := job.NextRun()
		if err != nil {
			return nil, fmt.Errorf("get next run for task %s | %w", job.Name(), err)
		}

		if j, ok := jobsRecurring[taskUID]; ok {
			lastRun, err := job.LastRun()
			if err != nil {
				return nil, fmt.Errorf("get last run for task %s | %w", job.Name(), err)
			}

			stats = append(stats, Stat{
				TaskUID:    j.task.UID,
				Name:       j.task.Name,
				Status:     j.status,
				NextRun:    nextRun,
				LastStatus: j.lastStatus,
				LastRun:    lastRun,
				LastError:  j.lastError,
				Interval:   j.interval,
			})
		}
	}

	slices.SortFunc(stats, func(a, b Stat) int { return int(a.NextRun.Sub(b.NextRun).Nanoseconds()) })

	return stats, nil
}

func (m *manager) wrap(task Task) func(context.Context) {
	return func(ctx context.Context) {
		m.mu.Lock()
		info, ok := m.jobs[task.UID()]
		if !ok {
			// Should not be possible
			m.mu.Unlock()
			zap.S().Errorf("Task %s not found during execution", task.Name())
			return
		}

		var user *model.User
		if len(info.users) > 0 {
			user = &info.users[0]
			info.users = info.users[1:]
		}
		info.status = Running

		m.jobs[task.UID()] = info
		m.mu.Unlock()

		// Run task
		start := time.Now()
		err := task.Func()(ctx, user)
		end := time.Now()

		// Save result
		userID := 0
		if user != nil {
			userID = user.ID
		}
		result := model.TaskSuccess
		if err != nil {
			result = model.TaskFailed
		}

		taskDB := &model.Task{
			UID:      task.UID(),
			UserID:   userID,
			RunAt:    time.Now(),
			Result:   result,
			Error:    err,
			Duration: end.Sub(start),
		}

		if errDB := m.repo.CreateRun(ctx, taskDB); errDB != nil {
			zap.S().Errorf("Failed to save recurring task result in database %+v | %v", *taskDB, err)
		}

		m.mu.Lock()
		defer m.mu.Unlock()

		info = m.jobs[task.UID()]
		info.status = Waiting
		info.lastStatus = result
		info.lastError = err
		m.jobs[task.UID()] = info
	}
}
