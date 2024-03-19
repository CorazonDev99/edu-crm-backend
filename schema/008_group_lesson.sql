-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE edu_group ALTER COLUMN edu_days TYPE VARCHAR(64),
  ALTER COLUMN edu_days SET DEFAULT '';
ALTER TABLE edu_group ALTER COLUMN lesson_start_time TYPE VARCHAR(64),
   ALTER COLUMN edu_days SET DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
