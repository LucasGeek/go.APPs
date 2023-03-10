package main

import "github.com/gofiber/fiber/v2"

func main() {
	server := fiber.New()

	server.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	server.Listen(":3000")
}
