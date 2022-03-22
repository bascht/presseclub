package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New(fiber.Config{})

	cacheDir, ok := os.LookupEnv("CACHE_DIR")
	if !ok {
		cacheDir = "/tmp"
	}

	log.Info("Using cache directory ", cacheDir)
	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html; charset=utf-8")

		return c.SendString("<h1 style='font-family: monospace'>ⴽ Ohai</h1>")
	})

	app.Get("/lies/*", func(c *fiber.Ctx) error {

		c.Set("Content-Type", "text/html")

		url := strings.TrimPrefix(c.OriginalURL(), "/lies/")
		urlHash := md5.Sum([]byte(url))
		cacheKey := hex.EncodeToString(urlHash[:])
		cachePath := filepath.Join(cacheDir, "presseclub."+cacheKey+".html")

		liesLogger := log.WithFields(log.Fields{"url": url, "cacheKey": cacheKey, "cachePath": cachePath})

		log.Info("Dealing with Request for ", url)

		if _, err := os.Stat(cachePath); err == nil {
			html, err := ioutil.ReadFile(cachePath)
			liesLogger.Info("Extracting HTML from cache")

			if err != nil {
				liesLogger.Error("Could not extract cache")

			}
			liesLogger.Info("Sending HTML from Cache")
			return c.SendString(string(html))

		} else if errors.Is(err, os.ErrNotExist) {
			liesLogger.Info("Downloading URL from source")
			command := "/usr/src/app/node_modules/single-file/cli/single-file"
			args := []string{
				"--browser-executable-path",
				"/usr/bin/chromium-browser",
				"--output-directory",
				"./../../../out/",
				"--browser-cookies-file",
				"/tmp/cookies.txt",
				"--browser-args", "[\"--no-sandbox\"]",
				"--dump-content",
				url,
			}
			html, err := exec.Command(command, args...).Output()

			if err != nil {
				return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
			}

			file, err := os.Create(cachePath)
			if err != nil {
				liesLogger.Error("Could open cache file")
			}

			file.Write(html)
			if err != nil {
				liesLogger.Error("Could not write cache to file")
			}

			return c.SendString(string(html))

		} else {
			liesLogger.Error("We fucked up big time. Could not read or write from cache directory")
			return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
		}
	})

	log.Fatal(app.Listen(":3000"))
}
