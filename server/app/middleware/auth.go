package middleware

import (
	"fiber-vue/app/utils/response"
	"github.com/gofiber/fiber/v2"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//	获取请求头中的token
		tokenString := c.Get("Authorization")
		//	验证token
		//	如果token无效
		//	返回错误信息
		//	如果token有效
		//	继续执行
		if tokenString == "" {
			return response.FailWithDetailed(nil, "未登录", c)
		}
		return c.Next()
	}
}
