package main

import (
	"go-typescript/database"
	"go-typescript/routes"
	"go-typescript/secret"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
	secret.LoadEnv("main", true)
	//DESC - Connect to database
	database.Conn.Connect()

	app := fiber.New(fiber.Config{})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use("storage", filesystem.New(filesystem.Config{
		Root: http.Dir("./storage"),
	}))

	//ANCHOR - Setup routes
	routes.Setup(app)

	//ANCHOR - Migrate and seed database (for development)
	app.Get("/mig", func(c *fiber.Ctx) error {
		database.Conn.DropSchema("public")
		database.Conn.Migrate()
		database.Conn.Seed()
		return c.SendString("public dropped and created")
	})

	app.Listen(":8000")
}
