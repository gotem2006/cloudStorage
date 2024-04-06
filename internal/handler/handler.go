package handler

import (
	"cloudstorage/internal/handler/storage"
	"cloudstorage/internal/service"

	"github.com/gofiber/fiber/v2"
)



type StorageHandler interface{
	PostFile(c *fiber.Ctx) error
	PostDirectory(c *fiber.Ctx) error
	GetFiles(c *fiber.Ctx) error
}

type Handler struct{
	Storage StorageHandler
}


func NewHandlers(services service.Service) *Handler{
	storageHandler := storage.NewStorageHandler(services.Storage)
	return &Handler{
		Storage: storageHandler,
	}
}