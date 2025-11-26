package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
		http.HandleFunc("/user", createUser)

	})

	log.Print("Listening on port 8443")
	if err := http.ListenAndServeTLS(":8443", "localhost.pem", "localhost-key.pem", nil); err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name string `json:"name"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sendJSONResponse(w, user)
}
func sendJSONResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
