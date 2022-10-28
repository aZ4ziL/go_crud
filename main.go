package main

import (
	"log"
	"net/http"

	"github.com/aZ4ziL/go_crud/handlers"
)

func main() {
	http.HandleFunc("/api/v1/data", handlers.DataAPIIndex)
	http.HandleFunc("/api/v1/data/post", handlers.DataAPIPost)
	http.HandleFunc("/api/v1/data/put", handlers.DataAPIPut)
	http.HandleFunc("/api/v1/data/delete", handlers.DataAPIDelete)

	log.Println("Server run...")
	http.ListenAndServe(":8000", nil)
}
