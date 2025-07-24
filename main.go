package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarxvrf/golang-projek/database"
	"github.com/tarxvrf/golang-projek/routers"
)

func main() {
	app := fiber.New()
	database.ConnectDB()
	routers.Handleroutes(app)
	app.Listen(":3000")

}
