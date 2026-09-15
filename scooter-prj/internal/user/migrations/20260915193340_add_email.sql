-- +goose Up
ALTER TABLE users 
ADD COLUMN email VARCHAR(255) NULL;

-- +goose Down
DROP COLUMN email;
