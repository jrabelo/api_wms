package UsuarioControllers

import (
	UsuarioModel "api_wms/src/models"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func Autenticar(ctx *fiber.Ctx) {
	ctx.Set("Access-Control-Allow-Origin", "*")
	ctx.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Set("Content-Type", "application/json")

	var body UsuarioModel.Request
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error:": "Não pode analisar json"})
		return
	}

	if len(strings.TrimSpace(body.Login)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Dados do Login está vazio!"})
		return
	}

	if len(strings.TrimSpace(body.Senha)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Dados da Senha está vazio!"})
		return
	}

	dados := UsuarioModel.Autenticar(body.Login, body.Senha)
	jwtSecret := []byte(os.Getenv("SECRET_KEY"))

	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["id"] = dados.ID
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	token, err := tk.SignedString(jwtSecret)
	if err != nil {
		log.Println(err)
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
