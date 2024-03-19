-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE auth_account ADD  CONSTRAINT auth_account FOREIGN KEY (role_id)
    REFERENCES auth_role(id) ON UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE auth_account ALTER COLUMN access_token SET DEFAULT '';
ALTER TABLE auth_account ALTER COLUMN refresh_token DROP NOT NULL ,
    ALTER  COLUMN refresh_token SET DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
