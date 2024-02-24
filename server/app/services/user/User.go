package user

import "github.com/gofiber/fiber/v2"

func CreateUser(username string, password string) fiber.Map {

	return fiber.Map{
		"message": "User created successfully" + " " + username,
		"status":  200,
	}
}

func FindUser() {

}
