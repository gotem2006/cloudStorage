package service

import (
	"cloudstorage/internal/service/storage"
	"github.com/gofiber/fiber/v2"
)





type StorageService interface{
	AddFile(c *fiber.Ctx) error
	AddDirectory(c *fiber.Ctx) error
	GetFiles(c *fiber.Ctx) error
}


type Service struct{
	Storage StorageService 
}



func NewServices() *Service{
	storageService := storage.NewStorageService()
	return &Service{
		Storage: storageService,
	}
}