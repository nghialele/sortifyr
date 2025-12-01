-- name: TaskRunCreate :one
INSERT INTO task_runs (task_uid, user_id, run_at, result, error, duration)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

