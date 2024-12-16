package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrazmee/go-crud2/controllers"
)

func TodoRoutes(app *fiber.App){
	app.Get("/todos", controllers.GetTodos)
	app.Get("/todos/:id", controllers.GetTodoById)
	app.Post("/todos", controllers.CreateTodo)
	app.Put("/todos/:id", controllers.UpdateTodo)
	app.Delete("/todos/:id", controllers.DeleteTodo)
}