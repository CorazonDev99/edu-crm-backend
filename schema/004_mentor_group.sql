-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION teacher_update_group_number() RETURNS TRIGGER LANGUAGE PLPGSQL AS
$$
  BEGIN
  --jsonni ichiga qo'yish kerak. number_of_groups deb. 
   UPDATE user SET  WHERE id = NEW.teacher_id;
    RETURN NEW;
  END;
$$;

CREATE TRIGGER teacher_number_of_groups_update_trigger AFTER INSERT ON group
FOR EACH ROW EXECUTE FUNCTION teacher_update_group_number();
-- +migrate StatementEnd
-- +migrate Down
