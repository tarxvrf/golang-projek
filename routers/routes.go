package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarxvrf/golang-projek/controller"
)

func Handleroutes(ctx *fiber.App) {
	ctx.Get("/", controller.HandleControl)
}
