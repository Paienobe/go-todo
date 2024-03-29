// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: tasks.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (id, name, isCompleted, user_id) 
VALUES($1, $2, $3, $4)
RETURNING id, name, iscompleted, user_id
`

type CreateTaskParams struct {
	ID          uuid.UUID
	Name        string
	Iscompleted bool
	UserID      uuid.UUID
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ID,
		arg.Name,
		arg.Iscompleted,
		arg.UserID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Iscompleted,
		&i.UserID,
	)
	return i, err
}

const deleteAllUserTasks = `-- name: DeleteAllUserTasks :exec
DELETE FROM tasks
WHERE user_id = $1
`

func (q *Queries) DeleteAllUserTasks(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteAllUserTasks, userID)
	return err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getUserTasks = `-- name: GetUserTasks :many
SELECT id, name, iscompleted, user_id FROM tasks WHERE user_id = $1
`

func (q *Queries) GetUserTasks(ctx context.Context, userID uuid.UUID) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getUserTasks, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Iscompleted,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const toggleTaskStatus = `-- name: ToggleTaskStatus :one
UPDATE tasks
SET isCompleted = $1
WHERE id = $2
RETURNING id, name, iscompleted, user_id
`

type ToggleTaskStatusParams struct {
	Iscompleted bool
	ID          uuid.UUID
}

func (q *Queries) ToggleTaskStatus(ctx context.Context, arg ToggleTaskStatusParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, toggleTaskStatus, arg.Iscompleted, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Iscompleted,
		&i.UserID,
	)
	return i, err
}
