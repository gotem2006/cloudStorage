package routes

import (
	"cloudstorage/internal/handler"
	"cloudstorage/internal/routes/storage"

	"github.com/gofiber/fiber/v2"
)




func InitRoutes(app *fiber.App, handler handler.Handler){
	storage.InitStorageRoutes(app, handler.Storage)
}