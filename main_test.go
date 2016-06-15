package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// type User struct {
// 	username      string
// 	password      string
// 	fact          []string
// 	factStripped  string
// 	facts         []string
// 	factsStripped string
// }

// Instantiate router
var router = buildRouter()

// Dummy user data
var newUser = struct {
	username      string
	password      string
	fact          []string
	factStripped  string
	facts         []string
	factsStripped string
}{
	"ian",
	"ian123",
	[]string{"日本語盛り上がりの"},
	"日本語盛上",
	[]string{"名称は、", "宇宙の膨張を発見した天文学者・エドウィン", "ハッブルに因む。"},
	"名称宇宙膨張発見天文学者因",
}

// Test cases
var testMatrix = []struct {
	method         string
	path           string
	body           string
	expectedStatus int
	expectedBody   string
}{
	{"GET", "/api/kanji", "", 403, ""},
	{"POST", "/api/signup", "", 201, ""},
	{"POST", "/", `{}`, 200, `{ss}`},
}

func TestAPI(t *testing.T) {

	// Test GET requests
	for _, tc := range testMatrix {

		// Build request
		var request *http.Request
		switch tc.method {
		case "GET":
			request, _ = http.NewRequest(
				tc.method,
				tc.path,
				nil,
			)
		case "POST":
			request, _ = http.NewRequest(
				tc.method,
				tc.path,
				strings.NewReader(tc.body),
			)
		}

		// Get response
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		// Check results
		fmt.Println(response.Body)
		var x map[string]interface{}
		var jsonResponse []byte
		decoder := json.NewDecoder(response.Body)
		decoder.Decode(&x)
		jsonResponse, _ = json.MarshalIndent(t, "", "    ")

		statusFail := response.Code != tc.expectedStatus
		bodyFail := string(jsonResponse) != tc.expectedBody
		if statusFail || bodyFail {
			t.Errorf(
				"\nMethod: %s, Path: %s\n"+
					"Expected status: %d\n"+
					"     Got status: %d\n"+
					"Expected body: %s\n"+
					"     Got body: %s\n",
				tc.method,
				tc.path,
				tc.expectedStatus,
				response.Code,
				tc.expectedBody,
				string(jsonResponse),
			)
		}

	}

}
