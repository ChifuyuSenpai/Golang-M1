package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		//////////////////
		// Premier try en front sans json
		//pid := os.Getpid()
		////response := fmt.Fprintf("Hostname: " + hostname + " | PID: " + string(rune(pid)) + " | Status: OK"), 2éme try
		//response := fmt.Sprintf("Time: %s | Hostname: %s | PID: %d | Status: OK", // \ IA
		//	time.Now().Format(time.DateTime),
		//	hostname,
		//	pid)
		//
		//w.Write([]byte(response))

		//w.Write([]byte(time.Now().Format(time.DateTime))) // 1er try
		//w.Write([]byte(" Status : OK"))
		//
		///////////////////

	})

	log.Print("Service UP » Listening on port 8443 !")
	if err := http.ListenAndServeTLS(":8443", "localhost.pem", "localhost-key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
