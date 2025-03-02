-- +goose Up
CREATE TABLE note (
    id BIGINT PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT,
    author BIGINT,
    dead_line TIMESTAMP,
    status BOOLEAN,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE persons (
    id BIGINT PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL
);

-- +goose Down
drop table note;
drop table persons;