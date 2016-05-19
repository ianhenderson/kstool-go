package main

import (
	// Internal libs
	"fmt"
	"net/http"
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
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler)
	router.HandleFunc("/{pageId}", handler)
	http.ListenAndServe(":8080", router)
}
