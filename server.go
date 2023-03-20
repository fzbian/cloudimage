package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	app := echo.New()

	app.Static("/i/", "images")
	app.Static("/", "index.html")
	app.GET("/health", Health)
	app.POST("/upload", Upload)

	err := app.Start(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func Health(c echo.Context) error {
	err := c.String(http.StatusOK, "Ok!")
	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}

func Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err.Error())
	}

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

	name := RandomName() + filepath.Ext(file.Filename)

	dst, err := os.Create("images/" + name)
	if err != nil {
		fmt.Println(err.Error())
	}

	if _, err = io.Copy(dst, open); err != nil {
		fmt.Println(err.Error())
	}

	return c.String(http.StatusOK, "File upload, "+name)
}

func RandomName() string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWYZ"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, 12)

	for i := range result {
		result[i] = charset[r.Intn(len(charset))]
	}

	return string(result)
}
