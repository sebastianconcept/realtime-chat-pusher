package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pusher/pusher-http-go"
)

func SetMessagesAPI(app *fiber.App, pusherClient pusher.Client) {

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		var data map[string]string

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient.Trigger("chat", "message", data)

		return c.JSON([]string{})
	})
}
