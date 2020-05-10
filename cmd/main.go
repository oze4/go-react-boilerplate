package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/template"
)

// PublicRoot defines the root path for static assets
type PublicRoot struct {
	path string
}

func (pr *PublicRoot) resolveFileName(n string) string {
	return fmt.Sprintf("%s/%s", strings.Trim(pr.path, "\\/"), strings.Trim(n, "\\/"))
}

func main() {
	app := fiber.New()
	staticDir := PublicRoot{path: "./public"}

	app.Static("/", staticDir.path)

	app.Use(helmet.New())
	app.Use(logger.New())

	app.Settings.TemplateEngine = template.Handlebars()

	app.Get("*", func(c *fiber.Ctx) {
		// Can also use `fiber.map{...}` instead of `map[string]interface{}{...}`
		data := map[string]interface{}{
			"message": "Hello, world!",
		}

		// This just shows how you can use Handlebars to render a template if needed.
		// In order to use with React, you would just place the built files into "./public".
		render := staticDir.resolveFileName("index.hbs")
		if err := c.Render(render, data); err != nil {
			c.Status(500).Send(err.Error())
		}
	})

	if err := app.Listen(3000); err != nil {
		log.Fatal(err)
	}
}
