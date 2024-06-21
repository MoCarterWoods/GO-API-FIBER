package main
import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	book.ID = len(books) + 1

	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdated := new(Book)

	if err := c.BodyParser(bookUpdated); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdated.Title
			books[i].Author = bookUpdated.Author
			return c.JSON(books[i])
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Book not found")
	}

	func deleteBook(c *fiber.Ctx) error {
		bookId, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		for i, book := range books {
			if book.ID == bookId {
				books = append(books[:i], books[i+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).SendString("Book not found")
	}