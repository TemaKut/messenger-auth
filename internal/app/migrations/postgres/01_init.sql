-- +goose Up
CREATE TABLE users (
    id      UUID            PRIMARY KEY,
    email   VARCHAR(255)    NOT NULL UNIQUE,
    data    JSONB           DEFAULT '{}'
);

CREATE INDEX users_email_b_tree_idx ON users (email);

-- +goose Down
DROP INDEX users_email_b_tree_idx;

DROP TABLE users;