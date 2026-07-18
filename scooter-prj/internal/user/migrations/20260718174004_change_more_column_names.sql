-- +goose Up
ALTER TABLE users RENAME last_name TO lastname;
ALTER TABLE users DROP COLUMN email;

-- +goose Down
ALTER TABLE users RENAME lastname TO last_name;
