package main

import (
	"golang-mod/database"
	"golang-mod/routers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	routers.Handleroutes(app)
	app.Listen(":3000")

}
