package main

import (
	"api_wms/src/routers"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	routers.LoadRouters(app)

	log.Fatal(app.Listen("3000"))
}
