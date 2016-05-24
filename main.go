package main

import (
	// Internal libs
	"fmt"
	"net/http"
	"os"
	// External libs
	"github.com/gorilla/mux"
)

func buildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler)
	router.HandleFunc("/{pageId}", handler)
	return router
}

func CreateServer(defaultPort string) error {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}
	router := buildRouter()
	fmt.Println("Listening on port", port)
	return http.ListenAndServe(":"+port, router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output := "Hi! You've requested: /" + vars["pageId"]
	fmt.Fprintf(w, output)
}

func main() {
	CreateServer("8000")
}
