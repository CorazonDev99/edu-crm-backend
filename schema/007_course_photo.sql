-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE course ADD COLUMN photo  VARCHAR(64) DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
