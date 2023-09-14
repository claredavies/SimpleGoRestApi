// main_test.go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

    // simulating a new HTTP GET request to the "/books" endpoint (nil means no body in request) (mock request)
	request := httptest.NewRequest(http.MethodGet, "/books", nil)
	// tracker used to analyse the response of the test (mock response)
	response := httptest.NewRecorder()
	// creating a new context with req and rec (mock)
	c := e.NewContext(request, response)

	//checking when we send a get request to /books
	 //1. that the response msg is http.StatusOK
	if assert.NoError(t, getBooks(c)) {
	    fmt.Printf("Status wanted: %v \n", http.StatusOK)
	    fmt.Printf("Returned Status: %v \n", response.Code)
		assert.Equal(t, http.StatusOK, response.Code)

		var responseBooks []book
// 		The code attempts to decode the JSON content from the HTTP response body
//      into the responseBooks variable, and halts the test with an error if decoding fails.
		if err := json.Unmarshal(response.Body.Bytes(), &responseBooks); err != nil {
			t.Fatalf("Failed reading response: %s", err)
		}
	    fmt.Printf("books: %v \n", books)
		assert.Equal(t, books, responseBooks)
	}
}