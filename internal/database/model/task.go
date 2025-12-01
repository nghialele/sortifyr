package model

import (
	"errors"
	"time"

	"github.com/topvennie/spotify_organizer/pkg/sqlc"
)

type TaskResult string

const (
	TaskSuccess TaskResult = "success"
	TaskFailed  TaskResult = "failed"
)

type Task struct {
	// Task result
	ID       int // ID of the task result
	UserID   int // ID of the user that started the task. 0 if it was scheduled
	RunAt    time.Time
	Result   TaskResult
	Error    error
	Duration time.Duration

	// Task fields
	UID    string // Identifier of the task
	Name   string
	Active bool
}

func TaskModel(task sqlc.Task, taskRun sqlc.TaskRun) *Task {
	userID := 0
	if taskRun.UserID.Valid {
		userID = int(taskRun.UserID.Int32)
	}
	var err error
	if taskRun.Error.Valid {
		err = errors.New(taskRun.Error.String)
	}

	return &Task{
		ID:       int(taskRun.ID),
		UserID:   userID,
		RunAt:    taskRun.RunAt.Time,
		Result:   TaskResult(taskRun.Result),
		Error:    err,
		Duration: time.Duration(taskRun.Duration),
		UID:      task.Uid,
		Name:     task.Name,
		Active:   task.Active,
	}
}

type TaskFilter struct {
	UserID  int
	TaskUID string
	Result  *TaskResult
	Limit   int
	Offset  int
}
