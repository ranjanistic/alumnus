package main

import (
  "github.com/joho/godotenv"
  "github.com/gofiber/fiber/v2"
  "alumnus/database"
  "os"
  "fmt"
)

func main() {
  godotenv.Load()
  connected,err := mongodb.Connect(os.Getenv("DBUSER"),os.Getenv("DBPASS"))
  if(err!=nil) {
    fmt.Println(err)
  } else if(connected){
    fmt.Println("Connected to database")
    app := fiber.New()
    app.Get("/login", func(c *fiber.Ctx) error {
      return c.SendString("Login post")
    })

    app.Get("/signup", func(c *fiber.Ctx) error {
      return c.SendString("signup post")
    })

    app.Use(func(c *fiber.Ctx) error {
      return c.SendStatus(404)
    })

    app.Listen("localhost:8000") 
  }
}