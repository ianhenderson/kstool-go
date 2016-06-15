package main

import (
	// Internal libs
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// External libs
	"github.com/gorilla/mux"
)

type testStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func buildRouter() *mux.Router {
	var router *mux.Router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler)
	router.HandleFunc("/{pageId}", handler)
	return router
}

func createServer(defaultPort string) error {
	var (
		port   string      = os.Getenv("PORT")
		router *mux.Router = buildRouter()
	)
	if len(port) == 0 {
		port = defaultPort
	}
	fmt.Println("Listening on port", port)
	return http.ListenAndServe(":"+port, router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		vars   map[string]string = mux.Vars(r)
		output string
	)

	switch r.Method {
	case "GET":
		fmt.Println("It's a GEEET")
		output = "Hi! You've requested: /" + vars["pageId"]
	case "POST":
		fmt.Println("It's a POOOOOST")
		var decoder = json.NewDecoder(r.Body)
		var t testStruct
		decoder.Decode(&t)
		fmt.Println("Username", t.Username)
		fmt.Println("Password", t.Password)
		var jsonResponse []byte
		jsonResponse, _ = json.Marshal(t)
		output = "Hi! You've requested: /" + vars["pageId"]
		output = output + "\n"
		output = output + string(jsonResponse[:])
	}

	fmt.Fprintf(w, output)
}

func main() {
	createServer("8000")
}
