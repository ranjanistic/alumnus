package main

import (
	"fmt"

	"github.com/ranjanistic/alumnus/config"
	"github.com/ranjanistic/alumnus/database"
	"github.com/ranjanistic/alumnus/router"
	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/template/html"
	"go.mongodb.org/mongo-driver/mongo"
)

func startUp(Users *mongo.Collection){
	app := fiber.New(fiber.Config{Views: html.New("./templates", ".html"), Prefork: config.Env.ENV == "production"})
	app.Use(csrf.New())
	app.Static("/", "./static")
	router.InitRoutes(app)
	app.Listen(":"+config.Env.PORT)
}

func main() {
	if config.Env.DEV && config.Env.ERR != nil {
		fmt.Printf("Local Environment: %s \n", config.Env.ERR)
	}
	fmt.Printf("Env: %s \n", config.Env.ENV)
	database.ConnectToDB(startUp)
}
