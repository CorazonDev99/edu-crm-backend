-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- CREATE TABLE IF NOT EXISTS group_rooms_enrollment (
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     group_room_id uuid NOT NULL ,
--     room_id uuid NOT NULL ,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (group_room_id) REFERENCES edu_group(id) ON DELETE CASCADE,
--     FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE
-- );
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down