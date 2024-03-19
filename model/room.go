package model

import "github.com/google/uuid"

type Room struct {
	ID          string `json:"id" db:"id"`
	Title       string `json:"title" db:"title" lenMin:"1" lenMax:"100" required:"true" regex:"login" `
	Description string `json:"description" db:"description" lenMin:"1" lenMax:"1024" required:"true" regex:"login"`
	RoomNumber  int    `json:"roomNumber" db:"room_number" amountMin:"0" amountMax:"100" required:"true" regex:"number"`
	OpenTime    string `json:"openTime" db:"open_time" required:"true" regex:"login"`
	CloseTime   string `json:"closeTime" db:"close_time" required:"true" regex:"login"`
	Total       int64  `json:"-" db:"total"`
}

type CreateRoom struct {
	Title       string `json:"title" db:"title" lenMin:"1" lenMax:"100" required:"true" regex:"login" `
	Description string `json:"description" db:"description" lenMin:"1" lenMax:"1024" required:"true" regex:"login"`
	RoomNumber  int    `json:"roomNumber" db:"room_number" amountMin:"0" amountMax:"100" required:"true" regex:"number"`
	OpenTime    string `json:"openTime" db:"open_time" required:"true" regex:"login"`
	CloseTime   string `json:"closeTime" db:"close_time" required:"true" regex:"login"`
}
type UpdateRoom struct {
	ID          uuid.UUID `json:"-" db:"id"`
	Title       string    `json:"title" db:"title" lenMin:"1" lenMax:"100" required:"true" regex:"login" `
	Description string    `json:"description" db:"description" lenMin:"1" lenMax:"1024" required:"true" regex:"login"`
	RoomNumber  int       `json:"roomNumber" db:"room_number" amountMin:"0" amountMax:"100" required:"true" regex:"number"`
	OpenTime    string    `json:"openTime" db:"open_time" required:"true" regex:"login"`
	CloseTime   string    `json:"closeTime" db:"close_time" required:"true" regex:"login"`
}
