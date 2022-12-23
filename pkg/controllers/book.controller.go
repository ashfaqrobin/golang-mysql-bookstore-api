package controllers

import (
	"net/http"

	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/models"
	"github.com/ashfaqrobin/golang-mysql-bookstore-api/pkg/storage"
	"github.com/gofiber/fiber/v2"
)

func CreateBooks(c *fiber.Ctx) error {
	book := models.Book{}

	err := c.BodyParser(&book)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't parse request body",
		})
	}

	err = storage.Repo.CreateBook(&book)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't create the book",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Book has been added",
	})
}

func GetBooks(c *fiber.Ctx) error {
	books := []models.Book{}

	err := storage.Repo.FindBooks(&books)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't find books",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    books,
	})
}

func GetBookById(c *fiber.Ctx) error {
	book := models.Book{}

	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "ID cann't be empty",
		})
	}

	err := storage.Repo.FindBookById(&book, id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't find book with id " + id,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    book,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	book := models.Book{}

	id := c.Params("id")

	bookBody := models.Book{}
	err := c.BodyParser(&bookBody)

	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't parse request body",
		})
	}

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "ID cann't be empty",
		})
	}

	err = storage.Repo.UpdateBook(&book, id, &bookBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't update book with id " + id,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	book := models.Book{}

	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "ID cann't be empty",
		})
	}

	err := storage.Repo.FindBookById(&book, id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't find book with id " + id,
		})
	}

	err = storage.Repo.DeleteBook(&book, id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": "Couldn't delete books",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Book has been deleted",
	})
}
