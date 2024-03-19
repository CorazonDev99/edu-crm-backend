-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE role_permission_enrollment ADD  FOREIGN KEY (role_id)
    REFERENCES auth_role(id) ON UPDATE CASCADE ON DELETE CASCADE;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
