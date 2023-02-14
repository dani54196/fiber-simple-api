package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	// enviromet variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// fiber
	app := fiber.New()
	// middlewares
	app.Use(logger.New())
	// index route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// group routes
	app.Use(requestid.New())
	userGroup := app.Group("/users")
	userGroup.Get("", handleUser)
	userGroup.Post("", handleCreateUser)
	// get port from env
	port := os.Getenv("PORT")
	// Listen
	app.Listen(":" + port)
	// fmt.Println(port)
}
