-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS crm_user(
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    full_name VARCHAR(64) NOT NULL,
    birthday_date TIMESTAMP NOT NULL,
    added_date TIMESTAMP NOT NULL,
    role VARCHAR(8) NOT NULL,
    phone_number VARCHAR(16) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    photo VARCHAR(32)  NULL,
    extra_data JSON NULL,
    -- comment TEXT NULL,
    -- tag varchar(32) NULL,
    -- learner_id INTEGER NULL,
    -- email VARCHAR(32) NULL,
    -- telegram VARCHAR(32) NULL,
    -- location VARCHAR(64) NULL,
    -- passport_image VARCHAR(64) NULL,
    -- password VARCHAR(32) NOT NULL,
    -- discord VARCHAR(32) NULL,
    -- parents_phone varchar(16)[] NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- CREATE TABLE IF NOT EXISTS card(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     title varchar(64) NOT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL
-- );
-- CREATE TABLE IF NOT EXISTS lid (
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     full_name varchar(64) NOT NULL,
--     phone_number varchar(16) NOT NULL,
--     location varchar(64) NOT NULL,
--     comment TEXT NULL,
--     role VARCHAR(8) NOT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL
-- );
-- CREATE TABLE IF NOT EXISTS card_lid_enrollment(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     card_id uuid NOT NULL,
--     lid_id uuid NOT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (card_id) REFERENCES card(id),
--     FOREIGN KEY (lid_id) REFERENCES lid(id)
-- );
-- CREATE TABLE IF NOT EXISTS learner_payment(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     learner_id uuid NOT NULL,
--     amount INTEGER NOT NULL,
--     date TIMESTAMP NOT NULL,
--     comment TEXT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (learner_id) REFERENCES learner(id)
-- );
-- CREATE TABLE IF NOT EXISTS edu_group_schedule(
--     id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     group_id UUID NOT NULL,
--     learner_id UUID NOT NULL,
--     date TIMESTAMP NOT NULL,
--     absent BOOLEAN DEFAULT FALSE,
--     homework BOOLEAN DEFAULT FALSE,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (group_id) REFERENCES edu_group(id),
--     FOREIGN KEY (learner_id) REFERENCES crm_user(id)
-- );
-- CREATE TABLE IF NOT EXISTS learner_attendance (
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     group_schedule_id UUID NOT NULL,
--     learner_id UUID NOT NULL,
--     absent BOOLEAN DEFAULT FALSE,
--     homework BOOLEAN DEFAULT FALSE,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (group_schedule_id) REFERENCES group_schedule(id),
--     FOREIGN KEY (learner_id) REFERENCES learner(id)
-- );
-- CREATE TABLE IF NOT EXISTS expense(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     title VARCHAR(64) NOT NULL,
--     date TIMESTAMP NOT NULL DEFAULT NOW(),
--     catalog VARCHAR(32) NOT NULl,
--     taker VARCHAR(64) NOT NULL,
--     amount INTEGER NOT NULL,
--     payment_type VARCHAR(16) NOT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL
-- );
-- CREATE TABLE IF NOT EXISTS course_learner_enrollment(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     course_id UUID NOT NULL,
--     learner_id UUID NOT NULL,
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (course_id) REFERENCES course(id),
--     FOREIGN KEY (learner_id) REFERENCES learner(id)
-- );
-- CREATE TABLE IF NOT EXISTS lesson(
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     course_id UUID NOT NULL,
--     title TEXT NOT NULL,
--     homework TEXT NULL DEFAULT '',
--     resource TEXT NULL DEFAULT '',
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL,
--     FOREIGN KEY (course_id) REFERENCES course(id)
-- );
-- CREATE TABLE IF NOT EXISTS (
--     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
--     created_at TIMESTAMP DEFAULT (NOW()),
--     updated_at TIMESTAMP NULL,
--     deleted_at TIMESTAMP NULL
-- );
