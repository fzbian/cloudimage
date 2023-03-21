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

	api := app.Group("/api")

	app.Static("/i/", "images")
	api.GET("/health", controllers.Health)
	api.POST("/upload", controllers.Upload)

	err = app.Start(":" + os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
