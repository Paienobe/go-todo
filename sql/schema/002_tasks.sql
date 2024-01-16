-- +goose Up
CREATE TABLE tasks (
    id UUID PRIMARY KEY NOT NULL,
    name TEXT NOT NULL,
    isCompleted BOOLEAN NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE 
);
-- +goose Down
DROP TABLE tasks; 