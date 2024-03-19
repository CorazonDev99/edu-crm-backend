-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE edu_group ALTER COLUMN room TYPE VARCHAR(64),
  ALTER COLUMN room SET DEFAULT '';
ALTER TABLE edu_group ALTER COLUMN price TYPE INTEGER,
   ALTER COLUMN edu_days SET DEFAULT 0;
ALTER TABLE edu_group ALTER COLUMN price TYPE INTEGER,
   ALTER COLUMN edu_days SET DEFAULT 0;
ALTER TABLE edu_group ALTER COLUMN comment TYPE TEXT,
   ALTER COLUMN comment SET DEFAULT '';
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
