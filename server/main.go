package main

import (
	"fiber-vue/app/database"
	"fiber-vue/app/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

// @title Fiber Admin App
// @version 1.0
// @description This is an API for Fiber Admin Application

// @contact.name hxl
// @contact.email 52553624@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {

	database.ConnectDB()

	app := fiber.New()

	router.Initalize(app)

	log.Fatal(app.Listen(":" + getenv("PORT", ":3000")))
}
