package main

import (
	"cloudimage/controllers"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/i/", "images")
	app.Static("/", "index.html")
	app.GET("/health", controllers.Health)
	app.POST("/upload", controllers.Upload)

	err := app.Start(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
