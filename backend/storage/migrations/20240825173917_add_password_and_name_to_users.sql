-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN username VARCHAR(50) NOT NULL,
    ADD COLUMN password VARCHAR(255) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN username,
    DROP COLUMN password;
-- +goose StatementEnd
