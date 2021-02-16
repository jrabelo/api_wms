package Controller

import (
	Model "api_wms/src/models"
	"log"
	"strings"

	"github.com/gofiber/fiber"
)

func CarregaTodosPedidos(ctx *fiber.Ctx) {
	ctx.Set("Access-Control-Allow-Headers", "*")
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Set("Access-Control-Allow-Methods", "*")

	var request Model.RequestApp
	err := ctx.BodyParser(&request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": "Não pode analisar json"})
		return
	}

	if len(strings.TrimSpace(request.ID_USER)) == 0 || len(strings.TrimSpace(request.ID_EMPRESA)) == 0 || len(strings.TrimSpace(request.JWT)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Parametros inválido!"})
		return
	}

	pedidos := Model.CarregaTodosPedidos(request.ID_EMPRESA)
	if pedidos == nil {
		log.Fatal("Pedidos: ", pedidos)
	}

	if err := ctx.JSON(pedidos); err != nil {
		ctx.Next(err)
	}

}
