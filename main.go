package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	////"opendavinci/routes"
	"opendavinci/database"
	"opendavinci/models"
)
import _ "embed"

//go:embed openapi.json
var oasJSON string

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	app := fiber.New(fiberConf())

	app.Get("/api/courses", func(c *fiber.Ctx) error {
		db, err := database.OpenDBConnection()
		if err != nil {
			// Return status 500 and database connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		courses := []models.Course{}
		query := `SELECT * FROM courses_v`
		err = db.Select(&courses, query)
		if err != nil {
			// Return status 500 and database connection error.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		return c.JSON(courses)
	})
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString(oasJSON)
	})

	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server failure: %v", err)
	}
}

func fiberConf() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
