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
	login := ctx.FormValue("login")
	senha := ctx.FormValue("senha")

	if len(strings.TrimSpace(login)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Campo Login está vazio!"})
		return
	}

	if len(strings.TrimSpace(senha)) == 0 {
		ctx.Status(200).JSON(fiber.Map{"msg": "Campo Senha está vazio!"})
		return
	}

	dados := UsuarioModel.Autenticar(login, senha)
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
