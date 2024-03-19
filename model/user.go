package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID              `json:"id"  db:"id"`
	FullName      string                 `json:"fullName" db:"full_name"`
	BirthdayDate  string                 `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber   string                 `json:"phoneNumber" db:"phone_number"`
	RoleID        uuid.UUID              `json:"roleID" db:"role_id"`
	Role          string                 `json:"role" db:"-"`
	Photo         string                 `json:"photo" db:"photo"`
	PhotoLink     string                 `json:"photoLink" db:"-"`
	AddedDate     string                 `json:"addedDate" db:"added_date"`
	ExtraDataByte []byte                 `json:"-"  db:"extra_data"`
	ExtraDataJSON map[string]interface{} `json:"extraData"  db:"-"`
	Total         int64                  `json:"-" db:"total"`
}

type CreateUser struct {
	FullName     string    `json:"fullName" db:"full_name" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate time.Time `json:"birthdayDate" db:"birthday_date"`
	PhoneNumber  string    `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	Password     string    `json:"password" db:"password" required:"true" lenMin:"0" lenMax:"64" regex:"password"`
	RoleID       uuid.UUID `json:"roleID" db:"role_id"`
	Photo        string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate    time.Time `json:"addedDate" db:"added_date" required:"true" `
	ExtraData    map[string]interface {
	} `json:"extraData"  db:"-"`
	ExtraDataJSON []byte `json:"-"  db:"extra_data"`
}
type UpdateUser struct {
	ID            uuid.UUID              `json:"-"  db:"id"`
	FullName      string                 `json:"fullName" db:"full_name"  required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	BirthdayDate  time.Time              `json:"birthdayDate" db:"birthday_date" required:"true" lenMin:"0" lenMax:"32"`
	PhoneNumber   string                 `json:"phoneNumber" db:"phone_number" required:"true" lenMin:"0" lenMax:"16" regex:"phone"`
	Password      string                 `json:"password" db:"password" required:"true" lenMin:"0" lenMax:"32" regex:"password"`
	RoleID        uuid.UUID              `json:"roleID" db:"role_id"`
	Photo         string                 `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	AddedDate     time.Time              `json:"addedDate" db:"added_date" required:"true" `
	ExtraData     map[string]interface{} `json:"extraDta" db:"-"`
	ExtraDataJSON []byte                 `json:"-"  db:"extra_data" required:"true"`
}
type SignInUser struct {
	PhoneNumber string `json:"phoneNumber" db:"phone_number" default:"+998901234567" required:"true"  lenMin:"0" lenMax:"16" regex:"phone"`
	Password    string `json:"password" db:"password" default:"EduCRM$007Boss" required:"true" lenMin:"0" lenMax:"16" regex:"password"`
}
type SignInUserResponse struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Role uuid.UUID `json:"role" db:"role_id"`
}

type UserPassword struct {
	Password string `json:"password"`
}
