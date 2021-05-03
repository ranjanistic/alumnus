package main

import (
	"fmt"

	"github.com/ranjanistic/alumnus/config"
	"github.com/ranjanistic/alumnus/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	if config.Env.ERR != nil {
		fmt.Printf("Local Environment: %s \n", config.Env.ERR)
	}
	fmt.Printf("Env: %s \n", config.Env.ENV)
	database.ConnectToDB(func(Users *mongo.Collection) {
		engine := html.New("./templates", ".html")
		app := fiber.New(fiber.Config{Views: engine, Prefork: config.Env.ENV == "production"})
		app.Static("/", "./static")
		auth := app.Group("/auth")
		auth.Get("/login",func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{"status":200})
		})
		app.Get("/", func(c *fiber.Ctx) error {
			return c.Render("index", fiber.Map{
				"title": config.Env.APPNAME,
			})
		})
		app.Listen(":3000")
	})
}
