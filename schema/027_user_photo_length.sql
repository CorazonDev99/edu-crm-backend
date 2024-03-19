-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE crm_user ALTER COLUMN photo TYPE VARCHAR(64);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
