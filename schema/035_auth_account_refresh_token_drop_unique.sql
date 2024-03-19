-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
ALTER TABLE auth_account DROP CONSTRAINT auth_account_refresh_token_key;
-- +migrate StatementEnd
-- +migrate Down
