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

	var body Model.RequestApp
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": "Não pode analisar json"})
		return
	}

	if len(strings.TrimSpace(body.ID_USER)) == 0 || len(strings.TrimSpace(body.ID_EMPRESA)) == 0 || len(strings.TrimSpace(body.JWT)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Parametros inválido!"})
		return
	}

	dados := Model.CarregaTodosPedidos(body.ID_EMPRESA)
	if dados == nil {
		log.Fatal("dados: ", dados)
	}

	if err := ctx.JSON(dados); err != nil {
		ctx.Next(err)
	}

}
