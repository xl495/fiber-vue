package controller

import (
	"fiber-vue/app/services/auth"
	"fiber-vue/app/services/user"
	"fiber-vue/app/utils/response"
	"github.com/gofiber/fiber/v2"
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

// BadRequestErrorResponse 定义一个通用的请求错误响应结构体
// swagger:model BadRequestErrorResponse
type BadRequestErrorResponse struct {
	// 错误信息
	Error string `json:"error"`
	// 具体错误详情（可选）
	Details map[string]interface{} `json:"details,omitempty"`
}

// InternalServerErrorResponse 定义一个通用的服务器内部错误响应结构体
// swagger:model InternalServerErrorResponse
type InternalServerErrorResponse struct {
	// 错误信息
	Error string `json:"error"`
	// 具体错误详情（可选）
	Details map[string]interface{} `json:"details,omitempty"`
}

// GetUser @Summary 获取用户
//
// @Produce json
// @Success 200 {object} string "成功"
// @Failure 400 {object} string "请求错误"
// @Failure 500 {object} string "内部错误"
// @Router /user [get]
func GetUser(ctx *fiber.Ctx) error {
	return ctx.JSON(user.GetAllUsers())
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建一个新的系统用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body UserReq true "用户信息"
// @Success 200 {object} model.User "成功创建用户并返回新创建的用户信息"
// @Failure 400 {object} BadRequestErrorResponse "请求错误，例如缺少必填字段或格式不正确"
// @Failure 500 {object} InternalServerErrorResponse "服务器内部错误"
// @Router /user [post]
func CreateUser(ctx *fiber.Ctx) error {
	var userReq UserReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}

	// BadRequestErrorResponse
	if userReq.Username == "" || userReq.Password == "" {
		return response.FailWithDetailed(nil, "用户名或密码不能为空", ctx)
	}

	return response.OkWithData(user.CreateUser(userReq.Username, userReq.Password), ctx)
}

// UserLogin 登录
// @Summary 登录
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body UserLoginReq true "用户信息"
// @Success 200 {object} model.User "成功创建用户并返回新创建的用户信息"
// @Failure 400 {object} BadRequestErrorResponse "请求错误，例如缺少必填字段或格式不正确"
// @Failure 500 {object} InternalServerErrorResponse "服务器内部错误"
// @Router /login [post]
func UserLogin(ctx *fiber.Ctx) error {
	var userReq UserLoginReq
	if err := ctx.BodyParser(&userReq); err != nil {
		return err
	}

	// BadRequestErrorResponse
	if userReq.Username == "" || userReq.Password == "" {
		return response.FailWithDetailed(nil, "用户名或密码不能为空", ctx)
	}

	return response.OkWithData(auth.UserLogin(userReq.Username, userReq.Password), ctx)

}
