-- name: CreateUser :one
INSERT INTO users (id, name, email, created_at, apikey)
VALUES($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING *;

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE apikey = $1;

-- name: GetUserByNameAndEmail :one
SELECT * FROM users WHERE name = $1 AND email = $2;