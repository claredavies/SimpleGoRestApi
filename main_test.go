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

func TestGetBooks_HappyPath(t *testing.T) {
    fmt.Println("--------TestGetBooks_HappyPath-----------")
	_, response, c := getMockRequestResponseContext(http.MethodGet, "/books")

	if assert.NoError(t, getBooks(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
		responseBooks := getMockResponseBooks(t, response)
		assert.Equal(t, books, responseBooks)
	}
}

func TestGetBookById_HappyPath(t *testing.T) {
    fmt.Println("--------TestGetBookById_HappyPath-----------")
    _, response, c := getMockRequestResponseContext(http.MethodGet, "/books")

    testBook := books[0]

    c.SetParamNames("id")
    c.SetParamValues(testBook.ID)

    if assert.NoError(t, bookById(c)) {
        assert.Equal(t, http.StatusOK, response.Code)

        responseBook := getMockResponseBook(t, response)

        assert.Equal(t, testBook, responseBook)
    }
}

func TestGetBookById_NoID(t *testing.T) {
    fmt.Println("--------TestGetBookById_NoID-----------")
    _, response, c := getMockRequestResponseContext(http.MethodGet, "/books")

    // NOT setting the 'id' parameter in the context

    if assert.NoError(t, bookById(c)) {
        assert.Equal(t, http.StatusBadRequest, response.Code)

        responseError := getMockResponseError(t, response)

        assert.Equal(t, ErrMsgParamIDRequired, responseError["message"])
    }
}

func TestGetBookById_IdDoesNotExist(t *testing.T) {
    fmt.Println("--------TestGetBookById_IdDoesNotExist-----------")
    _, response, c := getMockRequestResponseContext(http.MethodGet, "/books")

    //setting param id to non-existing value
    c.SetParamNames("id")
    c.SetParamValues("eee")

    if assert.NoError(t, bookById(c)) {
        assert.Equal(t, http.StatusNotFound, response.Code)

        responseError := getMockResponseError(t, response)

        assert.Equal(t, ErrMsgBookNotFound, responseError["message"])
    }
}

func getMockRequestResponseContext(method, path string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context){
    e := echo.New()
    SetupRoutes(e)

    request := httptest.NewRequest(method, path, nil)
    response := httptest.NewRecorder()
    c := e.NewContext(request, response)

    return e, response, c
}

func getMockResponseError(t *testing.T, response *httptest.ResponseRecorder) map[string]string {
    var responseError map[string]string
    if err := json.Unmarshal(response.Body.Bytes(), &responseError); err != nil {
        t.Fatalf("Failed reading response: %s", err)
    }
    return responseError
}

func getMockResponseBook(t *testing.T, response *httptest.ResponseRecorder) Book {
    var responseBook Book
    if err := json.Unmarshal(response.Body.Bytes(), &responseBook); err != nil {
        t.Fatalf("Failed reading response: %s", err)
    }
    return responseBook
}

func getMockResponseBooks(t *testing.T, response *httptest.ResponseRecorder) []Book {
     var responseBooks []Book
     if err := json.Unmarshal(response.Body.Bytes(), &responseBooks); err != nil {
        t.Fatalf("Failed reading response: %s", err)
     }
     return responseBooks
}
