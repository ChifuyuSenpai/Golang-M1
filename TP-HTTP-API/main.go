package main

import (
	//	"errors"
	"log"
	"net/http"
	//	"os"
	//	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Print("Service UP » Listening on port 8433 !")
	if err := http.ListenAndServeTLS(":8433", "localhost.pem", "localhost-key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
