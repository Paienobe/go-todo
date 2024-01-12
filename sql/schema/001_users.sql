-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE users;