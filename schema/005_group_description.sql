-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE edu_group ADD COLUMN description TEXT DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
