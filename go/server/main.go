package main

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{'health': 'ok'}")
}

func main() {
	fmt.Println("launch server")
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
