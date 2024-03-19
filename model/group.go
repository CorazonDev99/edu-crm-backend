package model

import (
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"  required:"true"`
	Description     string    `json:"description" db:"description"`
	CourseID        uuid.UUID `json:"courseID" db:"course_id"`
	TeacherID       uuid.UUID `json:"teacherID" db:"teacher_id"`
	EduDays         string    `json:"eduDays" db:"edu_days" required:"true"`
	RoomID          uuid.UUID `json:"roomId" db:"room_id"`
	RoomTitle       uuid.UUID `json:"room" db:"-"`
	Price           int       `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000"`
	Status          bool      `json:"status" db:"status"`
	LessonStartTime string    `json:"lessonStartTime" db:"lesson_start_time" `
	StartDate       string    `json:"startDate" db:"start_date"`
	EndDate         string    `json:"endDate" db:"end_date"`
	Comment         string    `json:"comment" db:"comment"`
	Total           int64     `json:"-" db:"total"`
}
type CreateGroup struct {
	Title           string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64" regex:"login"`
	Description     string    `json:"description" db:"description"  lenMin:"0" lenMax:"1024" regex:"login"`
	CourseID        uuid.UUID `json:"courseID" db:"course_id" required:"true"`
	TeacherID       uuid.UUID `json:"teacherID" db:"teacher_id" required:"true"`
	EduDays         string    `json:"eduDays" db:"edu_days" required:"true" lenMin:"0" lenMax:"32" regex:"login"`
	RoomID          uuid.UUID `json:"roomID" db:"room_id"`
	Price           int       `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000" regex:"login"`
	Status          bool      `json:"status" db:"status" `
	LessonStartTime string    `json:"lessonStartTime" db:"lesson_start_time"  lenMin:"0" lenMax:"16" regex:"login"`
	StartDate       time.Time `json:"startDate" db:"start_date" required:"true"`
	EndDate         time.Time `json:"endDate" db:"end_date" required:"true"`
	Comment         string    `json:"comment" db:"comment" lenMin:"0" lenMax:"2048" regex:"login"`
}
type UpdateGroup struct {
	ID              uuid.UUID `json:"-" db:"id"`
	Title           string    `json:"title" db:"title" required:"true" lenMin:"0" lenMax:"64"`
	Description     string    `json:"description" db:"description"  lenMin:"0" lenMax:"1024"`
	CourseID        uuid.UUID `json:"courseID" db:"course_id" required:"true"`
	TeacherID       uuid.UUID `json:"teacherID" db:"teacher_id" required:"true"`
	EduDays         string    `json:"eduDays" db:"edu_days" required:"true" lenMin:"0" lenMax:"32"`
	RoomID          uuid.UUID `json:"roomID" db:"room_id"`
	Price           int       `json:"price" db:"price"  required:"true" amountMin:"100000" amountMax:"100000000"`
	Status          bool      `json:"status" db:"status" `
	LessonStartTime string    `json:"lessonStartTime" db:"lesson_start_time"  lenMin:"0" lenMax:"16"`
	StartDate       time.Time `json:"startDate" db:"start_date" required:"true"`
	EndDate         time.Time `json:"endDate" db:"end_date" required:"true"`
	Comment         string    `json:"comment" db:"comment" lenMin:"0" lenMax:"2048"`
}
