package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	book "github.com/balickim/go-fiber/book"
)

// var mySigningKey = os.Get("JWT_KET")
var mySigningKey = []byte("supersecret")

func GenerateJWT(user string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Minute * 3).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v2/book", book.GetBooks)
	app.Get("/api/v2/book/:id", book.GetBook)
	app.Post("/api/v2/book", book.NewBook)
	app.Delete("/api/v2/book/:id", book.DeleteBook)
}

func main() {
    app := fiber.New()

	tokenString, err := GenerateJWT("Balu")
	if err != nil {
	fmt.Println("Error while generating token")
	}

	fmt.Println(tokenString)

    setupRoutes(app)

    app.Listen("127.0.0.1:3000")
}