-- name: CreateUser :exec
INSERT INTO users (id, name, email, created_at, password)
VALUES($1, $2, $3, $4, $5);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;