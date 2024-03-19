package model

import "github.com/google/uuid"

type Course struct {
	ID             uuid.UUID `json:"id" db:"id"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description" `
	Duration       string    `json:"duration" db:"duration"  `
	Status         bool      `json:"status" db:"status"`
	Price          int       `json:"price" db:"price"`
	Photo          string    `json:"photo" db:"photo"`
	PhotoLink      string    `json:"photoLink" db:"-"`
	LessonDuration string    `json:"lessonDuration" db:"lesson_duration"`
	Total          int64     `json:"-" db:"total"`
	//studentCount
}
type CreateCourse struct {
	Title          string `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Duration       string `json:"duration" db:"duration"  required:"true" lenMax:"64" regex:"login"`
	Description    string `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
	Status         bool   `json:"status" db:"status"`
	Photo          string `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Price          int    `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000"`
	LessonDuration string `json:"lessonDuration" db:"lesson_duration" required:"true"  lenMin:"0" lenMax:"16" regex:"login"`
}
type UpdateCourse struct {
	ID             uuid.UUID `json:"-" db:"id"`
	Title          string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Duration       string    `json:"duration" db:"duration"  required:"true" lenMax:"64" regex:"login"`
	Description    string    `json:"description" db:"description" lenMin:"0" lenMax:"1024" regex:"login"`
	Status         bool      `json:"status" db:"status"`
	Photo          string    `json:"photo" db:"photo" lenMin:"0" lenMax:"64" regex:"login"`
	Price          int       `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000"`
	LessonDuration string    `json:"lessonDuration" db:"lesson_duration" required:"true"  lenMin:"0" lenMax:"16" regex:"login"`
}
