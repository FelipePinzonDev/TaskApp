package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2"
)

type Todo struct{
	ID int `json:"id"`
	Completed bool `json:"completed"`
	Body	string `json:"body"`
}

//Have in mind that this & is refering to a memory space of something.
//And this * is a pointer that help u to have the value that is on that memory space

func main() {
	fmt.Println("Hello World's")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//Create task
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} //for todo is asign the space memory of Todo

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo) // Here u call for the value that is on that memory space

		return c.Status(201).JSON(todo)
	})

	//Update task
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error" : "Task not found"})
	})

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success" : true})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error" : "Task not found"})

	})

	app.Listen(":" + PORT)

}