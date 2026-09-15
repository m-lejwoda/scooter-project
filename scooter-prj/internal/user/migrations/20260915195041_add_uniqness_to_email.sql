-- +goose Up
ALTER TABLE users 
    ALTER COLUMN email TYPE VARCHAR(255),
    ALTER COLUMN email SET NOT NULL,
    ADD CONSTRAINT users_email_key UNIQUE (email),
    ADD CONSTRAINT users_email_check CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$');

-- +goose Down
ALTER TABLE users ALTER COLUMN email TYPE VARCHAR(255) NULL;
