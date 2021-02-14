package routers

import (
	Controller "api_wms/src/controllers"

	"github.com/gofiber/fiber"
)

func LoadRouters(app *fiber.App) {
	app.Post("/api/login", Controller.AutenticarUsuarios)
	app.Get("/api/pedings", Controller.CarregaTodosPedidos)

	app.Use("/api", func(ctx *fiber.Ctx) {
		ctx.Status(200).JSON(fiber.Map{"msg": "Não foi possível completar a solicitação!"})
	})

	app.Use("/", func(ctx *fiber.Ctx) {
		ctx.Status(200).JSON(fiber.Map{"msg": "Tudo funcionando..."})
	})
}
