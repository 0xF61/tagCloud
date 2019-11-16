package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tag", tagHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
