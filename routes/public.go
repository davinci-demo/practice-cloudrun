package routes

import (
	"github.com/gofiber/fiber/v2"
	"opendavinci/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/courses", controllers.GetCourses)              // get list of all courses
	route.Get("/course/:id", controllers.GetCourse)            // get one course by ID
	route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
