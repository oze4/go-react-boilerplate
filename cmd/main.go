package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/gofiber/helmet"
	"github.com/gofiber/logger"
	"github.com/gofiber/template"
	"github.com/joho/godotenv"
)

func main() {
	envConfig := dotEnvConfig{exitOnError: false}
	initDotEnv(envConfig)

	app := fiber.New()

	staticDir := PublicRoot{path: "./public"}
	app.Static("/", staticDir.path)

	app.Use(helmet.New())
	app.Use(logger.New())

	app.Settings.TemplateEngine = template.Handlebars()

	app.Get("*", func(c *fiber.Ctx) {
		// Can also use `map[string]interface{}{...}` instead of `fiber.map{...}`
		data := fiber.Map{
			"message": "Hello, world!",
		}

		// This just shows how you can use Handlebars to render a template if needed.
		// In order to use with React, you would just place the built files into "./public".
		render := staticDir.resolveFileName("index.hbs")
		if err := c.Render(render, data); err != nil {
			c.Status(500).Send(err.Error())
		}
	})

	port := "3030"
	if ep := os.Getenv("PORT"); ep != "" {
		port = ep
	}

	if err := app.Listen(port); err != nil {
		log.Fatal(err)
	}
}

// PublicRoot defines the root path for static assets
type PublicRoot struct {
	path string
}

func (pr *PublicRoot) resolveFileName(n string) string {
	return fmt.Sprintf("%s/%s", strings.Trim(pr.path, "\\/"), strings.Trim(n, "\\/"))
}

type dotEnvConfig struct {
	exitOnError bool
}

func initDotEnv(args dotEnvConfig) {
	if err := godotenv.Load(); err != nil {
		em := fmt.Sprintf("\n\n[ERROR][dotenv][PLEASE CONSIDER CREATING A .env FILE. YOU MAY ENCOUNTER BUGS OR EXPERIENCE ODD BEHAVIOR WITHOUT A .env FILE!] %s\n\n", err.Error())
		if args.exitOnError {
			log.Fatal(em)
		}
		fmt.Println(em)
	}
}
