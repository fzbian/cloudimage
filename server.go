package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func main() {
	port := 3000

	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	api := app.Group("/api")

	app.Get("/", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Hello world")
		if err != nil {
			return err
		}
		return nil
	})

	api.Get("/health", func(ctx *fiber.Ctx) error {
		err := ctx.SendString("Ok!")
		if err != nil {
			return err
		}
		return nil
	})

	fmt.Printf("Server on port %d", port)
	fiberPort := strconv.Itoa(port)
	err := app.Listen(":" + fiberPort)
	if err != nil {
		fmt.Println(err.Error())
	}
}
