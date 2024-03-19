-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE crm_user
    ADD COLUMN role_id UUID;
ALTER TABLE crm_user
    DROP COLUMN role;
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
