package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func handleHelloWorld(c *fiber.Ctx) error {
	fmt.Println("Sending Hello, World!")
	return c.SendString("Hello, World!")

}

func handleParams(c *fiber.Ctx) error {
	return c.SendString("Hello, " + c.Params("name"))
}

func handleAdd(c *fiber.Ctx) error {
	// atoi(III): convert  ASCII to integer from the Great C

	num1, err := strconv.Atoi(c.Params("num1"))
	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.SendString("Error Converting" + c.Params("num1") + "=>" + err.Error())
	}

	num2, err := strconv.Atoi(c.Params("num2"))

	if err != nil {
		c.SendStatus(fiber.StatusBadRequest)
		return c.SendString("Error Converting" + c.Params("num2") + "=>" + err.Error())
	}

	result := num1 + num2

	return c.SendString(strconv.Itoa(result))
}

// Have an endpoint to show Hello World

func main() {
	fmt.Println("Hello, World!")
	app := fiber.New()
	app.Get("/", handleHelloWorld)
	app.Get("/hello/:name", handleParams)
	app.Get("/add/:num1/:num2", handleAdd)
	app.Listen(":3000")
}
