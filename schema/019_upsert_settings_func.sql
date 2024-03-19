-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
-- +migrate StatementBegin
CREATE OR REPLACE  FUNCTION upsert_settings(company_title_var VARCHAR(64),
                                            company_logo_var VARCHAR(64), system_enter_logo_var VARCHAR(64) ,
                                            open_date_var TIMESTAMP, company_phone_var VARCHAR(64), site_color_var VARCHAR
        (16), instruction_file_var VARCHAR(64)) RETURNS VOID LANGUAGE PLPGSQL AS
$$
BEGIN
    IF  EXISTS (SELECT id  FROM settings  WHERE deleted_at IS NULL)
    THEN
        UPDATE settings SET company_title=company_title_var,
                            company_logo=company_logo_var,
                            system_enter_logo=system_enter_logo_var,open_date=open_date_var,company_phone=company_phone_var,site_color=site_color_var,instruction_file=instruction_file_var,updated_at = NOW()WHERE deleted_at IS NULL;
    ELSE
        INSERT INTO settings (company_title,company_logo,system_enter_logo,
                              open_date,company_phone,site_color,
                              instruction_file) VALUES (company_title_var, company_logo_var, system_enter_logo_var, open_date_var, company_phone_var, site_color_var, instruction_file_var);
    END IF;
END
$$
-- +migrate StatementEnd
-- +migrate Down
