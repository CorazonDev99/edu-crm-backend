-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied


CREATE TABLE IF NOT EXISTS board (
    id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    title varchar(64) NOT NULL,
    created_at TIMESTAMP DEFAULT (NOW()),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL
);
CREATE TABLE IF NOT EXISTS list (
                                     id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                                     board_id uuid NOT NULL,
                                     title varchar(64) NOT NULL,
                                     created_at TIMESTAMP DEFAULT (NOW()),
                                     updated_at TIMESTAMP NULL,
                                     deleted_at TIMESTAMP NULL,
                                     FOREIGN KEY (board_id) REFERENCES board(id)
);
CREATE TABLE IF NOT EXISTS lid (
                                   id uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
                                   list_id uuid NOT NULL,
                                   full_name varchar(64) NOT NULL,
                                   phone_number varchar(16) NOT NULL,
                                   location varchar(64) NOT NULL,
                                   comment TEXT NULL,
                                   created_at TIMESTAMP DEFAULT (NOW()),
                                   updated_at TIMESTAMP NULL,
                                   deleted_at TIMESTAMP NULL,
                                   FOREIGN KEY (list_id) REFERENCES list(id)
);
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
-- +migrate StatementEnd
-- +migrate Down
