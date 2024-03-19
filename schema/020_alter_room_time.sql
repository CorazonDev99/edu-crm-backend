-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE rooms ALTER COLUMN open_time TYPE VARCHAR(64);
ALTER TABLE rooms ALTER COLUMN close_time TYPE VARCHAR(64);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
