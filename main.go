package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c echo.Context) error {
	err := c.JSON(http.StatusOK, books)
    if err != nil {
        return err
    }

    return nil
}

func bookById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
            return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrMsgParamIDRequired})
    }

	book, err := getBookById(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": ErrMsgBookNotFound})
	}
	return c.JSON(http.StatusOK, book)
}

func getBookById(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New(ErrMsgBookNotFound)
}

func returnBook(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrMsgQueryIDRequired})
	}

	book, err := getBookById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": ErrMsgBookNotFound})
	}

	book.Quantity += 1
	return c.JSON(http.StatusOK, book)
}

func checkoutBook(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrMsgQueryIDRequired})
	}

	book, err := getBookById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": ErrMsgBookNotFound})
	}

	if book.Quantity <= 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": ErrMsgNoBooksRemaining})
	}

	book.Quantity -= 1
	return c.JSON(http.StatusOK, book)
}

func createBook(c echo.Context) error {
	var newBook Book
	if err := c.Bind(&newBook); err != nil {
		return err
	}
	books = append(books, newBook)
	return c.JSON(http.StatusCreated, newBook)
}

func SetupRoutes(e *echo.Echo) {
    e.GET("/books", getBooks)
    e.GET("/books/:id", bookById)
    e.POST("/books", createBook)
    e.PATCH("/checkout", checkoutBook)
    e.PATCH("/return", returnBook)
}

func main() {
	e := echo.New()

	SetupRoutes(e)
    e.Start("localhost:8080")
}