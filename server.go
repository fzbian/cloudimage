package main

import (
	"cloudimage/controllers"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app := echo.New()

	app.Static("/i/", "images")
	app.GET("/health", controllers.Health)
	app.POST("/upload", controllers.Upload)

	err = app.Start(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
