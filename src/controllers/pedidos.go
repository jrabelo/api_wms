package Controller

import (
	Model "api_wms/src/models"
	"strings"

	"github.com/gofiber/fiber"
)

func CarregaTodosPedidos(ctx *fiber.Ctx) {
	ctx.Set("Content-Type", "application/json")
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Set("Access-Control-Allow-Methods", "*")

	var body Model.RequestApp
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": "Não pode analisar json"})
		return
	}

	if body.ID == 0 || body.ID_FILIAL == 0 || len(strings.TrimSpace(body.Jwt)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Parametros inválido!"})
		return
	}

	response := fiber.Map{
		"status": "ok",
		"msg":    "Todos os pedidos...",
	}

	if err := ctx.JSON(response); err != nil {
		ctx.Next(err)
	}

}
