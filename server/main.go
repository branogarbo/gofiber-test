package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	server := fiber.New()

	server.Get('/', func(c *fiber.Ctx) {
		c.Send("bruh what are you doin")
	})

	server.Listen(3000)
}
