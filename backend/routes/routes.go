package routes

import (
	"github.com/gofiber/fiber/v2"
)

// ANCHOR - Function that combines parts of authority
func Setup(app fiber.Router) {
	api := app.Group("/api")
	User(api)

	// ANCHOR - No Authority
}
