-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE auth_role ADD UNIQUE (title);
ALTER TABLE auth_role ALTER COLUMN description  SET DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
