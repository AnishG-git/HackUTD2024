-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    email VARCHAR(255) NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    sdk_key TEXT NOT NULL UNIQUE
);

CREATE TABLE api_submissions (
    sdk_key TEXT NOT NULL,
    request_id TEXT NOT NULL,
    raw_req_id TEXT,
    raw_resp_id TEXT,
    req_in TIMESTAMP NOT NULL,
    req_out TIMESTAMP NOT NULL,
    endpoint TEXT NOT NULL,
    resp_states TEXT NOT NULL,
    method TEXT NOT NULL,
    PRIMARY KEY (sdk_key, request_id),
    FOREIGN KEY (sdk_key) REFERENCES users(sdk_key)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS api_submissions;
DROP TABLE IF EXISTS users;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
