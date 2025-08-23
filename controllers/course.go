package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"opendavinci/models"

	"opendavinci/database"
)

// GetCourses func gets all exists courses.
// @Description Get all exists courses.
// @Summary get all exists courses
// @Tags Courses
// @Accept json
// @Produce json
// @Success 200 {array} models.Course
// @Router /v1/courses [get]
func GetCourses(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all courses.
	courses, err := db.GetCourses()
	if err != nil {
		// Return, if courses not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"msg":     "courses were not found",
			"count":   0,
			"courses": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(courses),
		"courses": courses,
	})
}

// GetCourse func gets course by given ID or 404 error.
// @Description Get course by given ID.
// @Summary get course by given ID
// @Tags Course
// @Accept json
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} models.Course
// @Router /v1/course/{id} [get]
func GetCourse(c *fiber.Ctx) error {
	// Catch course ID from URL.
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get course by ID.
	course, err := db.GetCourse(id)
	if err != nil {
		// Return, if course not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "course with the given ID is not found",
			"course": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"course": course,
	})
}

// CreateCourse func for creates a new course.
// @Description Create a new course.
// @Summary create a new course
// @Tags Course
// @Accept json
// @Produce json
// @Param course_attrs body models.CourseAttrs true "Course JSON"
// @Success 200 {object} models.Course
// @Security ApiKeyAuth
// @Router /v1/course [post]
func CreateCourse(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current course.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Course struct
	course := &models.Course{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(course); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Course model.
	validate := NewValidator()

	// Set initialized default data for course:
	course.ID = uuid.New()
	course.Created = time.Now()
	////course.CourseStatus = 1 // 0 == draft, 1 == active

	// Validate course fields.
	if err := validate.Struct(course); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ValidatorErrors(err),
		})
	}

	if err := db.CreateCourse(course); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"course": course,
	})
}

// UpdateCourse func for updates course by given ID.
// @Description Update course.
// @Summary update course
// @Tags Course
// @Accept json
// @Produce json
// @Param id body string true "Course ID"
// @Param course_attrs body models.CourseAttrs true "Course JSON"
// @Success 201 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/course [put]
func UpdateCourse(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current course.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Course struct
	course := &models.Course{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(course); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if course with given ID does exist.
	foundedCourse, err := db.GetCourse(course.ID)
	if err != nil {
		// Return status 404 and course not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "course with this ID not found",
		})
	}

	// Set initialized default data for course:
	/////course.UpdatedAt = time.Now()

	// Create a new validator for a Course model.
	validate := NewValidator()

	// Validate course fields.
	if err := validate.Struct(course); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ValidatorErrors(err),
		})
	}

	// Update course by given ID.
	if err := db.UpdateCourse(foundedCourse.ID, course); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)
}

// DeleteCourse func for deletes course by given ID.
// @Description Delete course by given ID.
// @Summary delete course by given ID
// @Tags Course
// @Accept json
// @Produce json
// @Param id body string true "Course ID"
// @Success 204 {string} status "ok"
// @Security ApiKeyAuth
// @Router /v1/course [delete]
func DeleteCourse(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current course.
	expires := claims.Expires

	// Checking, if now time greater than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new Course struct
	course := &models.Course{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(course); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Course model.
	validate := NewValidator()

	// Validate only one course field ID.
	if err := validate.StructPartial(course, "id"); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if course with given ID does exist.
	foundedCourse, err := db.GetCourse(course.ID)
	if err != nil {
		// Return status 404 and course not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "course with this ID not found",
		})
	}

	// Delete course by given ID.
	if err := db.DeleteCourse(foundedCourse.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
