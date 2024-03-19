package model

import "github.com/google/uuid"

type Permission struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Tag         string    `json:"tag" db:"tag"`
	URL         string    `json:"url" db:"url"`
	Method      string    `json:"method" db:"method"`
	Total       int64     `json:"-" db:"total"`
}

type CreatePermission struct {
	Title       string `json:"title" db:"title" lenMin:"0" lenMax:"64" regex:"login" required:"true"`
	Description string `json:"description" db:"description" lenMin:"0" lenMax:"1024"`
	Tag         string `json:"tag" db:"tag" lenMin:"0" lenMax:"32" regex:"login" required:"true"`
	URL         string `json:"url" db:"url" lenMin:"0" lenMax:"64" regex:"login" required:"true"`
	Method      string `json:"method" db:"method" lenMin:"0" lenMax:"16" regex:"login" required:"true"`
}

type UpdatePermission struct {
	ID          uuid.UUID `json:"-" db:"id"`
	Title       string    `json:"title" db:"title" lenMin:"0" lenMax:"64" regex:"login" required:"true"`
	Description string    `json:"description" db:"description" lenMin:"0" lenMax:"1024"`
	Tag         string    `json:"tag" db:"tag" lenMin:"0" lenMax:"32" regex:"login" required:"true"`
	URL         string    `json:"url" db:"url" lenMin:"0" lenMax:"64" regex:"login" required:"true"`
	Method      string    `json:"method" db:"method" lenMin:"0" lenMax:"16" regex:"login" required:"true"`
}
