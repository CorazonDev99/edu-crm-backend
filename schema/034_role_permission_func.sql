-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION insert_role_permission() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
DECLARE
    permission_id UUID;
BEGIN
    FOR permission_id IN
        SELECT id from permission WHERE deleted_at IS NULL
        LOOP
            INSERT INTO role_permission_enrollment (role_id,permission_id,status) VALUES (NEW.id,permission_id,FALSE);
        END LOOP;
END;
$$;

CREATE TRIGGER insert_role_permission_trigger AFTER INSERT ON auth_role
    FOR EACH ROW
EXECUTE FUNCTION insert_role_permission();



CREATE OR REPLACE FUNCTION insert_permission_insert() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
DECLARE
    role_id UUID;
BEGIN
    FOR role_id IN
        SELECT id from auth_role WHERE deleted_at IS NULL
        LOOP
            INSERT INTO role_permission_enrollment (role_id,permission_id,status) VALUES (role_id,NEW.id,FALSE);
        END LOOP;
END;
$$;

CREATE TRIGGER insert_permission_insert_trigger AFTER INSERT ON permission
    FOR EACH ROW
EXECUTE FUNCTION insert_permission_insert();
-- +migrate StatementEnd
-- +migrate Down
