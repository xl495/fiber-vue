package controller

import (
	"fiber-vue/app/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type MenuReq struct {
	// 菜单名称
	// required: true
	// example: admin
	Name string `json:"name"`
}

func CreateMenu(ctx *fiber.Ctx) error {
	// 获取当前用户
	u := ctx.Locals("user")
	log.Info(u)

	var menu MenuReq
	if err := ctx.BodyParser(&menu); err != nil {
		return response.FailWithDetailed(nil, err.Error(), ctx)
	}

	return nil
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
	return nil
}
