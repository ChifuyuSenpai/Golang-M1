package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/c9s/goprocinfo/linux" // Library utilisée pour lire les infos CPU dans /proc/stat vu que Go n'a pas de package natif pour ça
)

// Depuis le usage repo GitHub
stat, err := linuxproc.ReadStat("/proc/stat")
if err != nil {
log.Fatal("stat read fail")
}

for _, s := range stat.CPUStats {
// s.User
// s.Nice
// s.System
// s.Idle
// s.IOWait
}


type HealthResponse struct {
	Time             string  `json:"time"`
	Hostname         string  `json:"hostname"`
	PID              int     `json:"pid"`
	Status           string  `json:"status"`
	GoVersion        string  `json:"go_version"`
	MemoryUsage      uint64  `json:"memory_usage"`
	MemoryAlloc      uint64  `json:"memory_alloc"`
	MemoryTotalAlloc uint64  `json:"memory_total_alloc"`
	CPUUsage         float64 `json:"cpu_usage"`
	CPUCores         int     `json:"cpu_cores"`
	// Stats pour calculer le usage CPU
	CPUUser          uint64 `json:"cpu_user"`
	CPUSystem        uint64 `json:"cpu_system"`
	CPUIdle          uint64 `json:"cpu_idle"`
}


func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		stat, err:= linuxproc.ReadStat("/proc/stat")
		if err != nil {
			log.Fatal("stat read fail")
		}
		var CPUUsage float64
		var CPUUser, CPUSystem, CPUIdle uint64

		if stat != nil && len(stat.CPUStatAll.CPUStats) > 0 {
			cpuStat := stat.CPUStatAll.CPUStats[0] // CPU total
			total := cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.Idle + cpuStat.IOWait + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal
			idle := cpuStat.Idle + cpuStat.IOWait
			CPUUsage = float64(total-idle) / float64(total) * 100.0
			CPUUser = cpuStat.User
			CPUSystem = cpuStat.System
			CPUIdle = cpuStat.Idle

			total := cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.Idle + cpuStat.IOWait + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal
			if total > 0 {
				CPUUsage = float64(total-idle) / float64(total) * 100.0
			} else {
				CPUUsage = "Error"
			}
		}

		response := HealthResponse{
			Time:             time.Now().Format(time.DateTime),
			Hostname:         hostname,
			PID:              os.Getpid(),
			Status:           "OK",
			GoVersion:        runtime.Version(),
			MemoryUsage:      m.Alloc / 1024 / 1024,
			MemoryAlloc:      m.TotalAlloc / 1024 / 1024,
			MemoryTotalAlloc: m.Sys / 1024 / 1024,
			CPUUsage: 	   CPUUsage,
			CPUCores:         runtime.NumCPU(),
			CPUUser:          CPUUser,
			CPUSystem:        CPUSystem,
			CPUIdle:          CPUIdle,
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
