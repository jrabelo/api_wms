package Controller

import (
	Model "api_wms/src/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func AutenticarUsuarios(ctx *fiber.Ctx) {
	ctx.Set("Access-Control-Allow-Headers", "*")
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Set("Access-Control-Allow-Methods", "*")

	var request Model.RequestLogin
	err := ctx.BodyParser(&request)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": "Não pode analisar json"})
		return
	}

	if len(strings.TrimSpace(request.Login)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Dados do Login está vazio!"})
		return
	}

	if len(strings.TrimSpace(request.Pass)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Dados da Senha está vazio!"})
		return
	}

	dados := Model.AutenticarUsuarios(request.Login, request.Pass)
	if dados == nil {
		log.Fatal("dados: ", dados)
	}

	jwtSecret := []byte(os.Getenv("SECRET_KEY"))

	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["id"] = dados.ID
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	token, err := tk.SignedString(jwtSecret)
	if err != nil {
		ctx.SendStatus(fiber.StatusInternalServerError)
	}

	response := fiber.Map{
		"status":       "ok",
		"id_user":      dados.ID,
		"name_user":    dados.Nome,
		"id_empresa":   dados.ID_FILIAL,
		"name_empresa": dados.Filial,
		"jwt":          token,
	}

	if err := ctx.JSON(response); err != nil {
		ctx.Next(err)
	}
}

func LogUsuario(ctx *fiber.Ctx) {
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

	AtualizaLog := Model.AtualizaLog(request.ID_EMPRESA, request.ID_USER)
	if AtualizaLog {
		response := fiber.Map{
			"msg": "Log do Usuario atualizado com sucesso!",
		}

		if err := ctx.JSON(response); err != nil {
			ctx.Next(err)
		}
	}
}
