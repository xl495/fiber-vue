package user

import (
	"errors"
	"fiber-vue/app/database"
	"fiber-vue/app/model"
	"fiber-vue/app/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllUsers get all users
func GetAllUsers() ([]model.User, error) {
	var users []model.User
	database.DB.Omit("password").Find(&users)
	return users, nil

}

func CreateUser(username string, password string) (fiber.Map, error) {

	pws, err := utils.PasswordHash(password)

	if err != nil {
		return nil, errors.New("密码加密错误")
	}

	// 查找用户名是否存在
	var existingUser model.User
	database.DB.Where("username = ?", username).First(&existingUser)

	if existingUser.ID != 0 {
		return nil, errors.New("用户名已存在")
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
		return nil, errors.New("创建用户失败")
	}

	return responseData, nil
}

func FindUser() {

}
