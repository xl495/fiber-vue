package user

import (
	"fiber-vue/app/database"
	"fiber-vue/app/model"
	"fiber-vue/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers get all users
func GetAllUsers() fiber.Map {
	var users []model.User
	database.DB.Find(&users)
	return fiber.Map{
		"code": fiber.StatusOK,
		"data": users,
	}

}

func CreateUser(username string, password string) fiber.Map {

	pws, err := utils.PasswordHash(password)

	if err != nil {
		return fiber.Map{
			"error": err,
		}
	}

	// 查找用户名是否存在
	var existingUser model.User
	database.DB.Where("username = ?", username).First(&existingUser)

	if existingUser.ID != 0 {
		return fiber.Map{
			"error": "用户名已存在",
			"code":  fiber.StatusBadRequest,
		}
	}

	newUser := &model.User{
		Username: username,
		Password: pws,
	}

	// 获取创建并返回用户数据 使用 bcrypt 加密密码
	dbUser := database.DB.Create(newUser)

	responseData := make(fiber.Map)
	responseData["ID"] = newUser.ID
	responseData["username"] = newUser.Username
	responseData["nickname"] = newUser.Nickname

	if dbUser.Error != nil {
		return fiber.Map{
			"error": dbUser.Error,
			"code":  fiber.StatusBadRequest,
		}
	}

	return fiber.Map{
		"data": responseData,
		"code": fiber.StatusOK,
	}
}

func FindUser() {

}
