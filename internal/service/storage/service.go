package storage

import (
	"fmt"
	"log"
	"os"
	_ "os"

	"github.com/gofiber/fiber/v2"
)



type storageService struct{

}


func NewStorageService() *storageService{
	return &storageService{}
}

var user = map[string]string{"1":"kirill"}

func (s *storageService) MakeFile(){
	
}

func (s *storageService) MakeDirectory(){

}


func (s *storageService) AddFile(c *fiber.Ctx) error{
	file, err := c.FormFile("file")
	if err != nil{
		return err
	}
	err = os.Mkdir(fmt.Sprintf("storage/%s", c.Params("user_id")), os.ModePerm)
	if err != nil{
		log.Println(err)
	}
	err = c.SaveFile(file, fmt.Sprintf("storage/%s/%s", c.Params("user_id"), file.Filename))
	return err
}


func (s *storageService) AddDirectory(c *fiber.Ctx) error{
	log.Println("i am works")
	form, err := c.MultipartForm()
	if err != nil{
		return err
	}
	err = os.Mkdir(fmt.Sprintf("storage/%s/%s", c.Params("user_id"), form.Value["name"][0]), os.ModePerm)
	files := form.File["file"]
	for _, file := range files{
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"])
		err = c.SaveFile(file, fmt.Sprintf("storage/%s/%s/%s", c.Params("user_id"), form.Value["name"][0], file.Filename))
		if err != nil{
			return err
		}
	}
	return nil
}


func (s *storageService) GetFiles(c *fiber.Ctx) error{
	files, err := os.ReadDir(fmt.Sprintf("storage/%s", c.Params("user_id")))
	if err !=nil{
		return err
	}
	response := ""
	for _, file := range files{
		response += file.Name() + "\n"
		if file.IsDir(){
			dirFiles, err := os.ReadDir(fmt.Sprintf("storage/%s/%s", c.Params("user_id"), file.Name()))
			if err != nil{
				return err
			}
			for _, dirFile := range dirFiles{
				response += dirFile.Name() + "\n"
			}
		}
		
	}
	return c.SendString(response)
}

func (s *storageService) UpdateFile(){

}

func (s *storageService) DeleteFile(){

}