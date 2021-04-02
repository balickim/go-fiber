package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/balickim/go-fiber/book"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v2/book", book.GetBooks)
	app.Get("/api/v2/book/:id", book.GetBook)
	app.Post("/api/v2/book", book.NewBook)
	app.Delete("/api/v2/book/:id", book.DeleteBook)
}

func main() {
    app := fiber.New()

    setupRoutes(app)

    app.Listen(":3000")
}