-- name: GetUserTasks :many
SELECT * FROM tasks WHERE user_id = $1;

-- name: CreateTask :one
INSERT INTO tasks (id, name, isCompleted, user_id) 
VALUES($1, $2, $3, $4)
RETURNING *;