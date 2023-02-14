package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type User struct {
	Id       string
	Name     string
	UserName string
}

func handleUser(c *fiber.Ctx) error {
	user := User{
		Name:     "pepe",
		UserName: "peperina",
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func handleCreateUser(c *fiber.Ctx) error {
	user := User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	user.Id = uuid.New().String()
	return c.Status(fiber.StatusOK).JSON(user)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userGroup := app.Group("/users")
	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)

	port := os.Getenv("PORT")

	// fmt.Println(port)
	app.Listen(":" + port)
}
