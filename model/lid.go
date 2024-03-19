package model

import "github.com/google/uuid"

type Board struct {
	ID    uuid.UUID `json:"id" db:"id"`
	Title string    `json:"title" db:"title"`
	List  []List    `json:"items"`
	Total int64     `json:"-" db:"total"`
}

type CreateBoard struct {
	Title string `json:"title" db:"title" regex:"login" lenMin:"3" lenMax:"64"`
}
type UpdateBoard struct {
	ID    uuid.UUID `json:"-" db:"id"`
	Title string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
}

type List struct {
	ID      uuid.UUID `json:"id" db:"id"`
	BoardID uuid.UUID `json:"boardID" db:"board_id"`
	Title   string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
	Total   int64     `json:"-" db:"total"`
	Lid     []Lid     `json:"lid"`
}

type CreateList struct {
	BoardID uuid.UUID `json:"boardID" db:"board_id"`
	Title   string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
}

type UpdateList struct {
	ID      uuid.UUID `json:"-" db:"id"`
	BoardID uuid.UUID `json:"boardID" db:"board_id"`
	Title   string    `json:"title" db:"title"  regex:"login" lenMin:"3" lenMax:"64"`
}

type Lid struct {
	ID          uuid.UUID `json:"id" db:"id"`
	ListID      uuid.UUID `json:"listID" db:"list_id"`
	FullName    string    `json:"fullName" db:"full_name"  regex:"login" lenMin:"3" lenMax:"64"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"  regex:"phone" lenMin:"3" lenMax:"16"`
	Location    string    `json:"location" db:"location"  regex:"login" lenMin:"3" lenMax:"64"`
	Comment     string    `json:"comment" db:"comment"  regex:"login" lenMin:"3" lenMax:"64"`
	Date        string    `json:"date" db:"created_at"`
	Total       int64     `json:"-" db:"total"`
}

type CreateLid struct {
	ListID      uuid.UUID `json:"listID" db:"list_id"`
	FullName    string    `json:"fullName" db:"full_name"  regex:"login" lenMin:"3" lenMax:"64"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"  regex:"phone" lenMin:"3" lenMax:"16"`
	Location    string    `json:"location" db:"location"  regex:"login" lenMin:"3" lenMax:"64"`
	Comment     string    `json:"comment" db:"comment"  regex:"login" lenMin:"3" lenMax:"1024"`
}

type UpdateLid struct {
	ID          uuid.UUID `json:"-" db:"id"`
	ListID      uuid.UUID `json:"listID" db:"list_id"`
	FullName    string    `json:"fullName" db:"full_name"  regex:"login" lenMin:"3" lenMax:"64"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number" regex:"phone" lenMin:"3" lenMax:"16"`
	Location    string    `json:"location" db:"location"  regex:"login" lenMin:"3" lenMax:"64"`
	Comment     string    `json:"comment" db:"comment"  regex:"login" lenMin:"3" lenMax:"1024"`
}

type MoveLid struct {
	From uuid.UUID `json:"from" db:"from"`
	To   uuid.UUID `json:"to" db:"to"`
}
type MoveList struct {
	From uuid.UUID `json:"from" db:"from"`
	To   uuid.UUID `json:"to" db:"to"`
}
