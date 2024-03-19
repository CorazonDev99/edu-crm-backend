-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE auth_role ADD document VARCHAR(64) NOT NULL default '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
