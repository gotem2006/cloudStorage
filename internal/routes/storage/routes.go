package storage

import (
	"cloudstorage/internal/handler"
	"github.com/gofiber/fiber/v2"
)



func InitStorageRoutes(app *fiber.App, handler handler.StorageHandler){
	app.Post("/:user_id", handler.PostFile)
	app.Post("/:user_id/directory", handler.PostDirectory)
	app.Get("/files/:user_id", handler.GetFiles)
}