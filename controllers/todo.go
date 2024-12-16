package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrazmee/go-crud2/database"
	"github.com/mrazmee/go-crud2/models"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.ToDo //struct ToDo
	database.DB.Find(&todos) //parameter disini adalah tempat untuk menampung hasil query
	return c.JSON(todos)
}

func GetTodoById(c *fiber.Ctx) error{
	id := c.Params("id")
	var todo models.ToDo

	result := database.DB.Find(&todo, id)

	if result.Error != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error{
	todo := new(models.ToDo)

	err := c.BodyParser(todo)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	database.DB.Create(&todo)
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error{
	id := c.Params("id")
	var todo models.ToDo

	result := database.DB.Find(&todo, id)
	if result.Error != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	if err := c.BodyParser(&todo); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":"Cannot Parse JSON",	
		})
	}

	database.DB.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error{
	id := c.Params("id")
	var todo models.ToDo

	result := database.DB.First(&todo, id)

	if result.Error != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	database.DB.Delete(&todo)
	return c.SendStatus(fiber.StatusNoContent)
}
