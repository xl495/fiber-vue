package auth

import (
	"errors"
	"fiber-vue/app/database"
	"fiber-vue/app/model"
	"fiber-vue/app/utils"
	"github.com/gofiber/fiber/v2"
)

func UserLogin(username string, password string) (fiber.Map, error) {
	//  1.校验用户是否存在
	//	2.校验密码
	//	3.生成token
	//	4.返回token
	var user model.User
	database.DB.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}

	isValid := utils.PasswordVerify(password, user.Password)

	if !isValid {
		return nil, errors.New("密码错误")
	}

	token, err := utils.GenerateToken(user.ID, user.Username)

	if err != nil {
		return nil, errors.New("token 生成错误")
	}

	return fiber.Map{
		"user":  user,
		"token": token,
	}, nil
}
