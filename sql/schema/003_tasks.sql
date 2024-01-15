-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    isCompleted BOOLEAN NOT NULL,
    user_id UUID REFERENCES users(id) NOT NULL
);
-- +goose Down
DROP TABLE tasks;