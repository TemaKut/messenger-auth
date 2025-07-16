-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    data JSONB
);

-- +goose Down
DROP TABLE users;