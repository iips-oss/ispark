package models

import (
	"time"

	"gorm.io/gorm"
)

// Track is a top-level grouping that activities belong to (e.g. "Personality
// Development", "Skill Building"). It is managed from the super admin Track
// Management screen. The number of activities in a track is derived from the
// Activity table (via Activity.TrackID), never stored here.
type Track struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(100);index;not null" json:"name"` // uniqueness enforced in the controller so a deleted track's name can be reused
	Description string         `gorm:"type:text" json:"description"`
	Status      string         `gorm:"type:varchar(20);not null;default:'Active'" json:"status"` // Active, Inactive
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// TotalActivities is computed per request, not persisted.
	TotalActivities int64 `gorm:"-" json:"total_activities"`
}

// TrackInput is the body accepted when creating or updating a track. Status is a
// pointer so an update can toggle it independently of the other fields.
type TrackInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Status      *string `json:"status"`
}
