package storage

import (
	"cloudstorage/internal/service"	
	"github.com/gofiber/fiber/v2"
)



type storageHandler struct{
	service service.StorageService
}


func NewStorageHandler(service service.StorageService) *storageHandler{
	return &storageHandler{service}
}


func (s *storageHandler) PostFile(c *fiber.Ctx) error{
	if err := s.service.AddFile(c); err != nil{
		return c.SendString(err.Error())
	}
	return c.SendString("Succes")
}


func (s *storageHandler) PostDirectory(c *fiber.Ctx) error{
	if err := s.service.AddDirectory(c); err != nil{
		return c.SendString(err.Error())
	}
	return c.SendString("Succes")
}


func (s *storageHandler) GetFiles(c *fiber.Ctx) error{
	return s.service.GetFiles(c)
}