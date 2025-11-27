package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthResponse struct {
	Time     string `json:"time"`
	Hostname string `json:"hostname"`
	PID      int    `json:"pid"`
	Status   string `json:"status"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		response := HealthResponse{
			Time:     time.Now().Format(time.DateTime),
			Hostname: hostname,
			PID:      os.Getpid(),
			Status:   "OK",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

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
