package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
)

func main() {
	app := fiber.New()

	PublicRoot := "./public"
	app.Static("/", PublicRoot)

	app.Use(helmet.New())
	app.Use(logger.New())

	app.Settings.TemplateEngine = template.Handlebars()

	app.Get("*", func(c *fiber.Ctx) {
		// Can also use `fiber.map{...}` instead of `map[string]interface{}{...}`
		data := map[string]interface{}{
			"message": "Hello, world!",
		}

		// This just shows how you can use Handlebars to render a template if needed, but
		// in order to use with React, you would just place the built files into './public'
		render := fmt.Sprintf("%s/index.hbs", PublicRoot)
		if err := c.Render(render, data); err != nil {
			c.Status(500).Send(err.Error())
		}
	})

	if err := app.Listen(3000); err != nil {
		log.Fatal(err)
	}
}
