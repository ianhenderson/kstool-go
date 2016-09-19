package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Helper function
func getCharAt(str string, i int) string {
    return string([]rune(str)[i])
}

// Instantiate router
var router = buildRouter()

// Dummy user data
var fakeUser = struct {
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

var badSigninInfo, _ = json.Marshal(
	map[string]string{
		"username": fakeUser.username,
		"password": "abcdefg",
	},
)

var newUserInfo, _ = json.Marshal(
	map[string]string{
		"username": fakeUser.username,
		"password": fakeUser.password,
	},
)

var expectedSigninResponse, _ = json.Marshal(
	map[string]interface{}{
		"id":   1,
		"name": fakeUser.username,
	},
)

var newWord, _ = json.Marshal(
	map[string][]string{
		"fact": fakeUser.fact,
	},
)

var newWords, _ = json.Marshal(
	map[string][]string{
		"fact": fakeUser.facts,
	},
)

// Test cases
var testMatrix = []struct {
	method         string
	path           string
	body           string
	expectedStatus int
	// expectedBody   string
	expectedBody   interface{}
}{
	// Get kanji w/out session
	{"GET", "/api/kanji", "", 403, ""},
	// Sign up
	{"POST", "/api/signup", string(newUserInfo), 201, string(expectedSigninResponse)},
	// Sign in w/ wrong info
	{"POST", "/api/login", string(badSigninInfo), 403, `{}`},
	// Sign in w/ correct info
	{"POST", "/api/login", string(newUserInfo), 200, string(expectedSigninResponse)},
	// Get kanji when list is empty
	{"GET", "/api/kanji", "", 404, ""},
	// Add word
	{"POST", "/api/facts", string(newWord), 201, ""},
	// TODO: check db integrity after adding one word
	// Get kanji when list is not empty
	{"GET", "/api/kanji", "", 200, getCharAt(fakeUser.factStripped, 0)}, // 日
	{"GET", "/api/kanji", "", 200, getCharAt(fakeUser.factStripped, 1)}, // 本
	{"GET", "/api/kanji", "", 200, getCharAt(fakeUser.factStripped, 2)}, // 語
	{"GET", "/api/kanji", "", 200, getCharAt(fakeUser.factStripped, 3)}, // 盛
	{"GET", "/api/kanji", "", 200, getCharAt(fakeUser.factStripped, 4)}, // 上
	// no more characters now
	{"GET", "/api/kanji", "", 404, ""},
	// Add multiple words
	{"POST", "/api/facts", string(newWords), 201, ""},
	// TODO: check db integrity after adding multiple words
	// Logout
	{"POST", "/api/logout", "", 200, ""},
	// Get kanji after logging out: not authorized
	{"GET", "/api/kanji", "", 403, ""},
}

func TestAPI(t *testing.T) {
	// Test GET requests
	for _, testCase := range testMatrix {

		// Build request
		var request *http.Request
		switch testCase.method {
		case "GET":
			request, _ = http.NewRequest(
				testCase.method,
				testCase.path,
				nil,
			)
		case "POST":
			request, _ = http.NewRequest(
				testCase.method,
				testCase.path,
				strings.NewReader(testCase.body),
			)
		}

		// Get response
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		// Check results
		var x map[string]interface{}
		json.NewDecoder(response.Body).Decode(&x)
		var jsonResponse, _ = json.Marshal(x)

		statusFail := response.Code != testCase.expectedStatus
		bodyFail := string(jsonResponse) != testCase.expectedBody
		if statusFail || bodyFail {
			t.Errorf(
				"\nMethod: %s, Path: %s\n"+
					"Expected status: %d\n"+
					"     Got status: %d\n"+
					"Expected body: %s\n"+
					"     Got body: %s\n",
				testCase.method,
				testCase.path,
				testCase.expectedStatus,
				response.Code,
				testCase.expectedBody,
				string(jsonResponse),
			)
		}

	}

}
