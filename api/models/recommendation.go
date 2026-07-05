package models

import (
	"time"
)

// Recommendation represents a course suggested by a mentor to a student
type Recommendation struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	StudentRollNo string    `gorm:"type:varchar(50);not null" json:"student_roll_no"`
	MentorID      string    `gorm:"type:varchar(50);not null" json:"mentor_id"`
	CourseName    string    `gorm:"type:varchar(100);not null" json:"course_name"`
	Message       string    `gorm:"type:text" json:"message"`
	CreatedAt     time.Time `json:"created_at"`
}
