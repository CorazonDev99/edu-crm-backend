-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS rooms (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    description TEXT NOT NULL,
    title VARCHAR(128) NOT NULL,
    room_number INT DEFAULT 0,
    open_time TIMESTAMP NOT NULL,
    close_time TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
