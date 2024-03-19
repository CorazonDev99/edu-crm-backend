package model

import "github.com/google/uuid"

type Role struct {
	ID           uuid.UUID `json:"id" db:"id" `
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Document     string    `json:"document" db:"document"`
	DocumentLink string    `json:"documentLink" db:"-"`
	Total        int64     `json:"-" db:"total"`
}
type CreateRole struct {
	Title       string `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64"`
	Description string `json:"description" db:"description" required:"false" lenMin:"0" lenMax:"1024"`
	Document    string `json:"document" db:"document"  lenMin:"0" lenMax:"64" regex:"login"`
}
type UpdateRole struct {
	ID          uuid.UUID `json:"-" db:"id"`
	Title       string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64"`
	Description string    `json:"description" db:"description" required:"false" lenMin:"0" lenMax:"1024"`
	Document    string    `json:"document" db:"document"  lenMin:"0" lenMax:"64" regex:"login"`
}
type AuthAccount struct {
	AccountID    uuid.UUID `json:"accountID" db:"account_id"`
	RoleID       uuid.UUID `json:"role" db:"role"`
	AccessToken  string    `json:"accessToken" db:"access_token"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
}
type CreateAuthAccount struct {
	AccountID    uuid.UUID `json:"accountID" db:"account_id"`
	RoleID       uuid.UUID `json:"roleId" db:"role_id"`
	AccessToken  string    `json:"accessToken" db:"access_token"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
}
type UpdateAuthAccount struct {
	ID           uuid.UUID `json:"id" db:"id"`
	AccountID    uuid.UUID `json:"accountID" db:"account_id"`
	RoleID       uuid.UUID `json:"roleId" db:"role_id"`
	AccessToken  string    `json:"accessToken" db:"access_token"`
	RefreshToken string    `json:"refreshToken" db:"refresh_token"`
}
