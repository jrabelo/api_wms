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
	app.Use(middleware.Recover())

	routers.LoadRouters(app)

	log.Fatal(app.Listen("8000"))
	//log.Fatal(app.Listen("192.168.0.169:8000"))
}
