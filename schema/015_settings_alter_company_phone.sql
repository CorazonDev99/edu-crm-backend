-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE settings ALTER  COLUMN company_phone TYPE VARCHAR(64);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
