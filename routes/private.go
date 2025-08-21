package routes

import (
	"github.com/gofiber/fiber/v2"
	"opendavinci/controllers"

)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/course", JWTProtected(), controllers.CreateCourse) // create a new course

	// Routes for PUT method:
	route.Put("/course", JWTProtected(), controllers.UpdateCourse) // update one course by ID

	// Routes for DELETE method:
	////route.Delete("/course", JWTProtected(), controllers.DeleteCourse) // delete one course by ID
}
