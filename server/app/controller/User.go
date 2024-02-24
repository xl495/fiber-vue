package controller

import (
	"fiber-vue/app/services/user"
	"github.com/gofiber/fiber/v2"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUser @Summary 获取用户
//
// @Produce json
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user [get]
func GetUser(ctx *fiber.Ctx) {
	return
}

// CreateUser @Summary 创建用户
//
// @Produce json
// @Param username body string true "账户名"
// @Param password body string true "密码"
// @Success 200 {object} model.User "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user [post]
func CreateUser(ctx *fiber.Ctx) error {
	var userReq UserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"data": user.CreateUser(userReq.Username, userReq.Password),
	})
}
