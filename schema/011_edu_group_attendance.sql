-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS edu_group_schedule(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    group_id UUID NOT NULL,
    lesson_title TEXT DEFAULT '',
    date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (group_id) REFERENCES edu_group(id)
);
CREATE TABLE IF NOT EXISTS edu_group_attendance(
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    learner_id UUID NOT NULL,
    group_schedule_id UUID NOT NULL,
    absent BOOLEAN DEFAULT FALSE,
    homework BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    FOREIGN KEY (group_schedule_id) REFERENCES edu_group_schedule(id),
    FOREIGN KEY (learner_id) REFERENCES crm_user(id)
);
-- +migrate Down
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
