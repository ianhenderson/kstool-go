package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router = buildRouter()

func TestIndex(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println("Response code:", response.Code)
	fmt.Println("Response body:", response.Body)
}

func TestBob(t *testing.T) {
	request, _ := http.NewRequest("GET", "/bob", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println("Response code:", response.Code)
	fmt.Println("Response body:", response.Body)
}

func TestJoe(t *testing.T) {
	request, _ := http.NewRequest("GET", "/joe", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println("Response code:", response.Code)
	fmt.Println("Response body:", response.Body)
}
