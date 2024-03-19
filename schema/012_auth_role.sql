-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS auth_account(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    account_id TEXT UNIQUE NOT NULL,
    refresh_token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
CREATE TABLE IF NOT EXISTS auth_role (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title VARCHAR(64) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
