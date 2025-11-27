package main

import (
	//	"errors"
	"log"
	"net/http"
	"time"
	//	"os"
	//	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(time.Now().Format(time.DateTime)))
		//		w.WriteHeader(os.)
	})

	log.Print("Service UP » Listening on port 8443 !")
	if err := http.ListenAndServeTLS(":8443", "localhost.pem", "localhost-key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
