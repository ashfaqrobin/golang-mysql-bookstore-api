package routes

import (
	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookStoreRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/books", controllers.CreateBooks)
	api.Get("/books", controllers.GetBooks)
	api.Get("/books/:id", controllers.GetBookById)
	api.Put("/books/:id", controllers.UpdateBook)
	api.Delete("/books/:id", controllers.DeleteBook)
}
