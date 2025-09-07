package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"opendavinci/models"
)

// CourseQueries struct for queries from Course model.
type CourseQueries struct {
	*sqlx.DB
}

// GetCourses method for getting all courses.
func (q *CourseQueries) GetCourses() ([]models.Course, error) {
	// Define courses variable.
	courses := []models.Course{}

	// Define query string.
	query := `SELECT * FROM courses_v`

	// Send query to database.
	err := q.Select(&courses, query)
	if err != nil {
		// Return empty object and error.
		return courses, err
	}

	// Return query result.
	return courses, nil
}

// GetCourse method for getting one course by given ID.
func (q *CourseQueries) GetCourse(id uuid.UUID) (models.Course, error) {
	// Define course variable.
	course := models.Course{}

	// Define query string.
	query := `SELECT * FROM courses_v WHERE id = $1`

	// Send query to database.
	err := q.Get(&course, query, id)
	if err != nil {
		// Return empty object and error.
		return course, err
	}

	// Return query result.
	return course, nil
}

// CreateCourse method for creating course by given json.
func (q *CourseQueries) CreateCourse(js string) error {
	query := `INSERT INTO courses (rawdata) VALUES ($1)`

	_, err := q.Exec(query, js)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// UpdateCourse method for updating course by given Course object.
func (q *CourseQueries) UpdateCourse(id uuid.UUID, b *models.Course) error {
	// Define query string.
	query := `UPDATE courses SET rawdata = $2 WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id, b.CourseID)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

// DeleteCourse method for delete course by given ID.
func (q *CourseQueries) DeleteCourse(id uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM courses WHERE id = $1`

	// Send query to database.
	_, err := q.Exec(query, id)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
