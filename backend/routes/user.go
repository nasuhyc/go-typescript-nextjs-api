package routes

import (
	c "go-typescript/controllers"

	"github.com/gofiber/fiber/v2"
)

// ANCHOR - Route for user
func User(api fiber.Router) {
	api = api.Group("/user")
	api.Get("/get/:id<int>", c.UserController{}.Get)
	api.Get("/getAll", c.UserController{}.Index)
	api.Post("/create", c.UserController{}.Store)
	api.Put("/update/:id<int>", c.UserController{}.Update)
	api.Delete("/delete/:id<int>", c.UserController{}.Delete)
}
