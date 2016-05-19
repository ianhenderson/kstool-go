package main

import (
	// Internal libs
	"fmt"
	"net/http"
	"os"
	// External libs
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	output := "Hi! You've requested: /" + vars["pageId"]
	fmt.Println(output)
	fmt.Fprintf(w, output)
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler)
	router.HandleFunc("/{pageId}", handler)
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, router)
}
