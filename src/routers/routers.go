package routers

import (
	UsuarioControllers "api_wms/src/controllers"

	"github.com/gofiber/fiber"
)

func LoadRouters(app *fiber.App) {
	app.Post("/api/login", UsuarioControllers.Autenticar)

	app.Use("/api", func(ctx *fiber.Ctx) {
		ctx.Status(200).JSON(fiber.Map{"msg": "Não foi possível completar a solicitação!"})
	})
}
