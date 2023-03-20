package controllers

import (
	"cloudimage/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

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

	name := utils.RandomName() + filepath.Ext(file.Filename)

	dst, err := os.Create("images/" + name)
	if err != nil {
		fmt.Println(err.Error())
	}

	if _, err = io.Copy(dst, open); err != nil {
		fmt.Println(err.Error())
	}

	return c.String(http.StatusOK, "File upload, "+name)
}
