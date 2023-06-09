package controllers

import (
	"cloudimage/utils"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Data struct {
	CreatedAt time.Time `yaml:"date"`
	IP        string    `yaml:"ip"`
}

func Upload(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	error := map[string]string{
		"error": "Se requiere una imagen para cargar.",
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, error)
	}

	NewFileName := utils.RandomName()
	ImageFileName := NewFileName + filepath.Ext(file.Filename)
	ImageInfoFile := NewFileName + ".yaml"

	open, err := file.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(open)

	FileInfo := Data{
		CreatedAt: time.Now(),
		IP:        c.RealIP(),
	}

	InfoFileStruct, err := os.Create("data/" + ImageInfoFile)
	if err != nil {
		panic(err)
	}
	defer func(InfoFile *os.File) {
		err := InfoFile.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(InfoFileStruct)

	encoder := yaml.NewEncoder(InfoFileStruct)
	err = encoder.Encode(FileInfo)
	if err != nil {
		panic(err)
	}

	dst, err := os.Create("images/" + ImageFileName)
	if err != nil {
		fmt.Println(err.Error())
	}

	if _, err = io.Copy(dst, open); err != nil {
		fmt.Println(err.Error())
	}

	data := map[string]string{
		"link": os.Getenv("DOMAIN") + "i/" + ImageFileName,
	}

	return c.JSON(http.StatusCreated, data)
}
