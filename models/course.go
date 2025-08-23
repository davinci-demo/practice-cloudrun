package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// Course struct to describe course object.
type Course struct {
	ID           uuid.UUID   `db:"id" json:"id" validate:"required,uuid"`
	Created      time.Time   `db:"created" json:"created"`
	CourseID     string      `db:"courseid" json:"courseId" validate:"required"`
	Title        string      `db:"title" json:"title" validate:"required,lte=255"`
	Instructor   string      `db:"instructor" json:"instructor" validate:"lte=255"`
	Descriptions string      `db:"description" json:"description" validate:"lte=255"`
	Image        string      `db:"image" json:"image" validate:"lte=255"`
	Published    time.Time   `db:"published" json:"published"`
	Updated      time.Time   `db:"updated" json:"updated"`
	CourseAttrs  CourseAttrs `db:"rawdata"`
}

// CourseAttrs struct to describe course attributes.
type CourseAttrs struct {
	Rawdata string
}

// Value make the CourseAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b *CourseAttrs) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// Scan make the CourseAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *CourseAttrs) Scan(value interface{}) error {
	j, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(j, &b)
}
