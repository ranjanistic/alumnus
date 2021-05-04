package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ranjanistic/alumnus/handler"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", handler.Root)
	app.Get("/alum", handler.Dash)
	app.Get("/callback", handler.CallbackHandler)
	auth := app.Group("/auth")
	auth.Get("/login", handler.Login)
	auth.Get("/signup", handler.Signup)
	auth.Get("/logout", handler.Logout)
	profile := app.Group("/profile")
	profile.Get("/",handler.Profile)
	profile.Get("/:username",handler.Profile)
	settings := app.Group("/settings")
	settings.Get("/:category",handler.Settings)
}
