package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
//     "github.com/labstack/echo/v4"
    "errors"
//     "fmt"
)

type book struct{
    // need uppercase variables so in go the variables are public but for json need them to be lowercase e.g. `json:"id"`
    ID string `json:"id"`
    Title string `json:"title"`
    Author string `json:"author"`
    Quantity int `json:"quantity"`
}

//slice of the book struct
var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

//gin.Context is all the information about the request and lets you get a response
func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
    id := c.Param("id")
    book, err := getBookById(id)

    if err != nil {
        //gin.H allows to write json wanted to return
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
        return
    }

    c.IndentedJSON(http.StatusOK, book)
}

func getBookById(id string) (*book, error) {
    for i, b := range books {
        if b.ID == id {
            return &books[i], nil
        }
    }
    return nil, errors.New("book not found")
}

func returnBook(c *gin.Context) {
    id, ok := c.GetQuery("id")

        if !ok {
            // StatusBadRequest as didn't pass the correct query parameter
            c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
            return
        }

        book, err := getBookById(id)
        if err != nil {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
            return
        }

        book.Quantity += 1
        c.IndentedJSON(http.StatusOK, book)
}
func checkoutBook(c *gin.Context) {
    id, ok := c.GetQuery("id")

    if !ok {
        // StatusBadRequest as didn't pass the correct query parameter
        c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
        return
    }

    book, err := getBookById(id)
    if err != nil {
        c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
        return
    }

    if book.Quantity <= 0 {
       c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "All books out on loan"})
       return
    }

    book.Quantity -= 1
    c.IndentedJSON(http.StatusOK, book)
}

func createBook(c *gin.Context) {
    var newBook book

    // trying to bind c json to newBook and if doesn't work return
    if err := c.BindJSON(&newBook); err != nil {
        return
    }
    books = append(books, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func main() {
    // router is from gin and allows the handling of different routes e.g. when go to localhost:8080/books it's going to call getBooks
    router := gin.Default()
    router.GET("/books", getBooks)
    router.GET("/books/:id", bookById)
    router.POST("/books", createBook)
    router.PATCH("/checkout", checkoutBook)
    router.PATCH("/return", returnBook)
    router.Run("localhost:8080")
}