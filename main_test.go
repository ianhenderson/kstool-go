package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type User struct {
	username      string
	password      string
	fact          []string
	factStripped  string
	facts         []string
	factsStripped string
}

var router = buildRouter()
var newUser = User{
	"ian",
	"ian123",
	[]string{"日本語盛り上がりの"},
	"日本語盛上",
	[]string{"名称は、", "宇宙の膨張を発見した天文学者・エドウィン", "ハッブルに因む。"},
	"名称宇宙膨張発見天文学者因",
}

var testMatrix = []struct {
	method         string
	path           string
	body           string
	expectedStatus int
	expectedBody   string
}{
	{"GET", "/", "", 200, ""},
	{"GET", "/bob", "", 200, ""},
	{"POST", "/", `{}`, 200, ""},
}

func TestAPI(t *testing.T) {
	// Test GET requests
	for _, tc := range testMatrix {
		request, _ := http.NewRequest(tc.method, tc.path, strings.NewReader(tc.body))
		response := httptest.NewRecorder()

		router.ServeHTTP(response, request)

		if response.Code != tc.expectedStatus {
			t.Errorf("Method: %s, Path: %s, Expected: %d, Got: %d", tc.method, tc.path, tc.expectedStatus, response.Code)
		}

	}

}

func TestBob(t *testing.T) {
	t.Skip()
	request, _ := http.NewRequest("GET", "/bob", nil)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println("Response code:", response.Code)
	fmt.Println("Response body:", response.Body)
}

func TestPost(t *testing.T) {
	t.Skip()
	body, _ := json.Marshal(newUser)
	request, _ := http.NewRequest("POST", "/bob", bytes.NewReader(body))
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println("Response code:", response.Code)
	fmt.Println("Response body:", response.Body)
}
