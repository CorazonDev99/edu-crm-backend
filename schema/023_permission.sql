-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS permission (
                                     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                                     title varchar(64) NOT NULL,
                                     description TEXT DEFAULT '',
                                     tag VARCHAR(128) DEFAULT  '',
                                     created_at TIMESTAMP DEFAULT (NOW()),
                                     updated_at TIMESTAMP NULL,
                                     deleted_at TIMESTAMP NULL
);
CREATE TABLE IF NOT EXISTS role_permission_enrollment (
                                     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                                     role_id uuid NOT NULL,
                                     permission_id uuid NOT NULL,
                                     status BOOLEAN  DEFAULT FALSE,
                                     created_at TIMESTAMP DEFAULT (NOW()),
                                     updated_at TIMESTAMP NULL,
                                     deleted_at TIMESTAMP NULL,
                                     FOREIGN KEY (permission_id) REFERENCES permission (id),
                                     FOREIGN KEY (role_id) REFERENCES  auth_role(id)
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
