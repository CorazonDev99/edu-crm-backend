package model

import (
	"time"

	"github.com/google/uuid"
)

type Settings struct {
	ID                  uuid.UUID `json:"-" db:"id"`
	CompanyTitle        string    `json:"companyTitle" db:"company_title" lenMin:"0" lenMax:"64" regex:"login"`
	CompanyLogo         string    `json:"-" db:"company_logo" required:"true" `
	CompanyLogoLink     string    `json:"companyLogoLink" db:"-" `
	SiteEnterLogo       string    `json:"-" db:"system_enter_logo" lenMin:"0" lenMax:"64"`
	SiteEnterLogoLink   string    `json:"siteEnterLogoLink" db:"-"`
	OpenDate            time.Time `json:"openDate" db:"open_date"`
	CompanyPhone        string    `json:"companyPhone" db:"company_phone" lenMin:"0" lenMax:"64" regex:"phone"`
	SiteColor           string    `json:"siteColor" db:"site_color" lenMin:"0" lenMax:"16" `
	InstructionFile     string    `json:"instructionFile" db:"instruction_file" lenMin:"0" lenMax:"64"  regex:"login"`
	InstructionFileLink string    `json:"instructionFileLink" db:"-"`
}
type CreateSettings struct {
	CompanyTitle    string    `json:"companyTitle" db:"company_title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	CompanyLogo     string    `json:"companyLogo" db:"company_logo" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	SiteEnterLogo   string    `json:"siteEnterLogo" db:"system_enter_logo" lenMin:"0" lenMax:"64" regex:"login"`
	OpenDate        time.Time `json:"openDate" db:"open_date"`
	CompanyPhone    string    `json:"companyPhone" db:"company_phone" lenMin:"0" lenMax:"64" regex:"phone"`
	SiteColor       string    `json:"siteColor" db:"site_color" lenMin:"0" lenMax:"16" regex:"hexColor"`
	InstructionFile string    `json:"instructionFile" db:"instruction_file" lenMin:"0" lenMax:"64"  regex:"login"`
}
