-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS course (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title VARCHAR(64) NOT NULL,
    description TEXT DEFAULT '',
    duration VARCHAR(8) NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    price INTEGER DEFAULT 0,
    lesson_duration  VARCHAR(32) NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
-- CREATE TABLE IF NOT EXISTS teacher_course_enrollment(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     course_id UUID NOT NULL,
--     teacher_id UUID NOT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (course_id) REFERENCES course(id),
--     FOREIGN KEY (teacher_id) REFERENCES crm_user(id)
-- );
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
