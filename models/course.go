package models

import (
	"time"

	"github.com/google/uuid"
)

// Course struct to describe course object.
type Course struct {
	ID           uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	Created      time.Time `db:"created" json:"created"`
	CourseID     string    `db:"courseid" json:"courseId" validate:"required"`
	Title        string    `db:"title" json:"title" validate:"required,lte=255"`
	Instructor   string    `db:"instructor" json:"instructor" validate:"lte=255"`
	Descriptions string    `db:"description" json:"description" validate:"lte=255"`
	Subject      string    `db:"subject" json:"subject" validate:"lte=255"`
	Image        string    `db:"image" json:"image" validate:"lte=255"`
	Published    string    `db:"published" json:"published"`
	Updated      string    `db:"updated" json:"updated"`
}
