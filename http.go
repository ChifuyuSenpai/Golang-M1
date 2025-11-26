package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/Hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")

	})

	log.Print("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Output: Hello, World! (at http://localhost:8080/Hello)
