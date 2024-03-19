-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE edu_group
    ADD COLUMN room_id UUID,
    ADD  CONSTRAINT room FOREIGN KEY (room_id)
    REFERENCES rooms(id) ON UPDATE CASCADE ON DELETE CASCADE;
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
