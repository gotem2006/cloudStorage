package main

import (
	"cloudstorage/internal/handler"
	"cloudstorage/internal/routes"
	"cloudstorage/internal/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	app := fiber.New()

	app.Use(cors.New())

	services := service.NewServices()
	handlers := handler.NewHandlers(*services)
	routes.InitRoutes(app, *handlers)

	err := app.Listen(":8080")
	if err != nil{
		log.Fatal(err)
	}
}
