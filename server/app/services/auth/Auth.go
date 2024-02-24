package auth

import (
	"fiber-vue/app/database"
	"fiber-vue/app/model"
	"fiber-vue/app/utils"
	"github.com/gofiber/fiber/v2"
)

func UserLogin(username string, password string) fiber.Map {
	//  1.校验用户是否存在
	//	2.校验密码
	//	3.生成token
	//	4.返回token
	var user model.User
	database.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return fiber.Map{
			"error": "用户不存在",
			"code":  fiber.StatusBadRequest,
			"data":  nil,
		}
	}

	isValid := utils.PasswordVerify(password, user.Password)

	if !isValid {
		return fiber.Map{
			"error": "密码错误",
			"code":  fiber.StatusBadRequest,
			"data":  nil,
		}
	}

	token, err := utils.GenerateToken(user.ID, user.Username)

	if err != nil {
		return fiber.Map{
			"error": err,
			"code":  fiber.StatusInternalServerError,
			"data":  nil,
		}
	}

	return fiber.Map{
		"code": fiber.StatusOK,
		"data": fiber.Map{
			"user":  user,
			"token": token,
		},
	}

}
