package models

import (
	"time"
)

type Certificate struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	StudentRollNo string    `gorm:"type:varchar(50);not null" json:"student_roll_no"`
	ActivityName  string    `gorm:"type:varchar(200);not null" json:"activity_name"`
	FileURL       string    `gorm:"type:text;not null" json:"file_url"`
	Status        string    `gorm:"type:varchar(20);default:'pending'" json:"status"` // pending, approved, rejected
	Feedback      string    `gorm:"type:text" json:"feedback"`
	VerifiedBy    string    `gorm:"type:varchar(50)" json:"verified_by"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}