package routers

import (
	"golang-mod/controller"

	"github.com/gofiber/fiber/v2"
)

func Handleroutes(ctx *fiber.App) {
	ctx.Get("/", controller.HandleControl)
}
