package controller

import (
	"fiber-vue/app/services/auth"
	"fiber-vue/app/services/user"
	"fiber-vue/app/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type UserReq struct {
	// 用户名
	// required: true
	// example: admin
	Username string `json:"username"`
	// 密码
	// required: true
	// example: 123456
	Password string `json:"password"`
}

type UserLoginReq struct {
	// 用户名
	// required: true
	// example: admin
	Username string `json:"username"`
	// 密码
	// required: true
	// example: 123456
	Password string `json:"password"`
}

// GetUser @Summary 获取用户
// @Produce json
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user [get]
// @schemes http
// @Security JWTAuthToken
// @SecurityDefinitions.JWTAuthToken
// @in header
// @name Authorization
func GetUser(ctx *fiber.Ctx) error {
	// 获取当前用户
	u := ctx.Locals("user")
	log.Info(u)

	result, err := user.GetAllUsers()

	if err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	return response.OkWithData(result, ctx)
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建一个新的系统用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body UserReq true "用户信息"
// @Success 200 {object} model.User "成功创建用户并返回新创建的用户信息"
// @Router /user [post]
func CreateUser(ctx *fiber.Ctx) error {
	var userReq UserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}

	if userReq.Username == "" || userReq.Password == "" {
		return response.FailWithDetailed(nil, "用户名或密码不能为空", ctx)
	}

	result, err := user.CreateUser(userReq.Username, userReq.Password)

	if err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	return response.OkWithData(result, ctx)
}

// UserLogin 登录
// @Summary 登录
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body UserLoginReq true "用户信息"
// @Success 200 {object} model.User "成功创建用户并返回新创建的用户信息"
// @Router /login [post]
func UserLogin(ctx *fiber.Ctx) error {
	var userReq UserLoginReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}

	if userReq.Username == "" || userReq.Password == "" {
		return response.FailWithDetailed(nil, "用户名或密码不能为空", ctx)
	}

	userMap, err := auth.UserLogin(userReq.Username, userReq.Password)

	if err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	return response.OkWithData(userMap, ctx)

}
