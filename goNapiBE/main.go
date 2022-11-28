package main

import (
	"goNapi/files"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {

	api := app.Group("/api")
	api.Get("/", files.GetApi)
	api.Get("/dirs", files.GetDirectory)
	api.Get("/files", files.GetFiles)
	api.Post("/exec", files.ExecQnapi)
}

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "goNapi project v1.0",
	})

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
