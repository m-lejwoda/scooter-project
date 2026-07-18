-- +goose Up
ALTER TABLE users RENAME COLUMN name TO username;
-- +goose Down
ALTER TABLE users RENAME COLUMN username TO name;
