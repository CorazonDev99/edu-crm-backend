-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE permission ADD COLUMN url  VARCHAR(64) DEFAULT '';
ALTER TABLE permission ADD COLUMN method  VARCHAR(16) DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
