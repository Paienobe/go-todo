-- name: GetUserTasks :many
SELECT * FROM tasks WHERE user_id = $1;

-- name: CreateTask :one
INSERT INTO tasks (id, name, isCompleted, user_id) 
VALUES($1, $2, $3, $4)
RETURNING *;

-- name: ToggleTaskStatus :one
UPDATE tasks
SET isCompleted = $1
WHERE id = $2
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;
