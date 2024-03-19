-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS edu_group (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title varchar(64) NOT NULL,
    course_id uuid  NULL,
    teacher_id uuid  NULL,
    edu_days varchar(32) NOT NULL,
    room varchar(16)  NULL,
    price INTEGER NOT NULL,
    lesson_start_time VARCHAR(16)  NULL,
    status BOOLEAN DEFAULT FALSE,
    -- start_date and end_date are used to calculate the duration of the course
    start_date TIMESTAMP  NULL,
    end_date TIMESTAMP  NULL,
    comment TEXT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (teacher_id) REFERENCES crm_user(id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);
CREATE TABLE IF NOT EXISTS edu_group_learner_enrollment (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    group_id uuid NOT NULL,
    learner_id uuid NOT NULL,
    start_date TIMESTAMP  DEFAULT NOW(),
    end_date TIMESTAMP  NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (group_id) REFERENCES edu_group(id),
    FOREIGN KEY (learner_id) REFERENCES crm_user(id)
);
CREATE TABLE IF NOT EXISTS edu_group_teacher_enrollment (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    group_id uuid NOT NULL,
    teacher_id uuid NOT NULL,
    start_date TIMESTAMP  DEFAULT NOW(),
    end_date TIMESTAMP  NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (group_id) REFERENCES edu_group(id),
    FOREIGN KEY (teacher_id) REFERENCES crm_user(id)
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
