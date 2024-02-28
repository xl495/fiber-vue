package router

import (
	"fiber-vue/app/controller"
	"fiber-vue/app/middleware"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {

	// swagger

	api := router.Group("/api")

	api.Use(swagger.New(swagger.Config{
		BasePath: "/api/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})

	api.Use(middleware.Json)

	api.Post("/user", controller.CreateUser)
	api.Get("/users", middleware.JwtMiddleware(), controller.GetUser)
	api.Post("/login", controller.UserLogin)

	needAuth := api.Group("/menu")

	// Menu API
	needAuth.Use(middleware.JwtMiddleware())

	needAuth.Post("/", controller.CreateMenu)
	needAuth.Put("/", controller.UpdateMenu)
	needAuth.Delete("/", controller.RemoveMenu)
	needAuth.Get("/", controller.GetMenu)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
