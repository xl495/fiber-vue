package controller

import (
	"fiber-vue/app/model"
	"fiber-vue/app/services"
	"fiber-vue/app/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type MenuReq struct {
	// 菜单名称
	// required: true
	// example: admin
	Name string `json:"name"`

	// 菜单路径
	// required: true
	// example: /admin
	Path string `json:"path"`

	// 父菜单ID
	// required: true
	// example: 0
	ParentId int `json:"parentId"`

	// 排序
	// required: true
	// example: 0
	Sort int `json:"sort"`

	// 组件
	// required: true
	// example: Admin
	Component string `json:"component"`

	// Mate
	// required: true
	Mate MateReq `json:"mate"`
}

type MateReq struct {
	// 是否隐藏
	// required: true
	// example: false
	Hidden bool `json:"hidden"`

	// 是否保持活动状态
	// required: true
	// example: false
	KeepAlive bool `json:"keepAlive"`

	// 图标
	// required: true
	// example: admin
	Icon string `json:"icon"`

	// 标题
	// required: true
	// example: 管理员
	Title string `json:"title"`
}

func CreateMenu(ctx *fiber.Ctx) error {
	// 获取当前用户
	u := ctx.Locals("user")
	log.Info(u)

	var menu model.Menu
	if err := ctx.BodyParser(&menu); err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	result, err := services.CreateMenu(menu)

	if err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	return response.OkWithData(&result, ctx)
}

func UpdateMenu(ctx *fiber.Ctx) error {
	// 获取当前用户
	u := ctx.Locals("user")
	log.Info(u)

	return nil
}

func RemoveMenu(ctx *fiber.Ctx) error {
	return nil
}

func GetMenu(ctx *fiber.Ctx) error {

	result, err := services.GetMenuList()

	if err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	return response.OkWithData(&result, ctx)
}
