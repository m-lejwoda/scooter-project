-- +goose Up
CREATE TABLE users (
  id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name VARCHAR(100) NOT NULL CHECK (length(trim(name)) >= 2),
  email VARCHAR(255) NOT NULL UNIQUE CHECK (length(trim(email)) >= 5),
  last_name VARCHAR(100),
  password VARCHAR(255) NOT NULL 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)
-- +goose Down
DROP TABLE users;
