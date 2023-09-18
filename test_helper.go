package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
    "fmt"
    "bytes"
	"github.com/labstack/echo/v4"
)

func getMockRequestResponseContext(method, path string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context){
    e := echo.New()
    SetupRoutes(e)

    request := httptest.NewRequest(method, path, nil)
    response := httptest.NewRecorder()
    c := e.NewContext(request, response)

    return e, response, c
}

func getMockRequestResponseContextWithQuery(method, path, queryKey, queryValue string) (*echo.Echo, *httptest.ResponseRecorder, echo.Context) {
    e := echo.New()
    SetupRoutes(e)

    request := httptest.NewRequest(method, path, nil)
    request.URL.RawQuery = fmt.Sprintf("%s=%s", queryKey, queryValue) // set the query parameter

    response := httptest.NewRecorder()
    c := e.NewContext(request, response)

    return e, response, c
}

func getMockRequestResponseContextWithRequestBody(method, path string, requestBody []byte) (*echo.Echo, *httptest.ResponseRecorder, echo.Context) {
    e := echo.New()
    SetupRoutes(e)

    // Create a request with the specified method, path, and request body
    request := httptest.NewRequest(method, path, bytes.NewReader(requestBody))
    request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set the content type to JSON

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