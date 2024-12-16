package main

import(
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/mrazmee/go-crud2/routes"
	"github.com/mrazmee/go-crud2/database"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.TodoRoutes(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	fmt.Println("Server running at http://localhost:3000")
}