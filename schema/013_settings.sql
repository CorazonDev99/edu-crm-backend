-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS settings (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    company_title VARCHAR(64) NOT NULL,
    company_logo VARCHAR(64)  NULL,
    system_enter_logo VARCHAR(64)  NULL,
    open_date TIMESTAMP  NULL,
    company_phone TEXT DEFAULT '',
    site_color VARCHAR(16) NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
