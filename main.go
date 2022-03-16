package main

import (
	"log"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new engine by passing the template folder
	// and template extension using <engine>.New(dir, ext string)

	app := fiber.New(fiber.Config{})

	app.Static("/downloads", "/downloads")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/lies/https://bascht.com")
	})

	app.Get("/lies/*", func(c *fiber.Ctx) error {

		url := strings.TrimPrefix(c.OriginalURL(), "/lies/")

		command := "/usr/src/app/node_modules/single-file/cli/single-file"
		args := []string{
			"--browser-executable-path",
			"/usr/bin/chromium-browser",
			"--output-directory",
			"./../../../out/",
			"--browser-args", "[\"--no-sandbox\"]",
			"--dump-content",
			url,
		}
		out, err := exec.Command(command, args...).Output()

		if err != nil {
		    return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
		}

		c.Set("Content-Type", "text/html")

		return c.SendString(string(out))
	})

	log.Fatal(app.Listen(":3000"))
}
