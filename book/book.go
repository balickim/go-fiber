package book

import "github.com/gofiber/fiber/v2"

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Book")
}

func NewBook(c *fiber.Ctx) error {
	return c.SendString("add new book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("delete book")
}