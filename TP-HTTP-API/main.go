package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/c9s/goprocinfo/linux"
)

var startTime = time.Now()

type HealthResponse struct {
	Time          string  `json:"time"`
	Hostname      string  `json:"hostname"`
	PID           int     `json:"pid"`
	Status        string  `json:"status"`
	GoVersion     string  `json:"go_version"`
	Uptime        string  `json:"uptime"`
	MemoryUsageMB uint64  `json:"memory_usage_mb"`
	MemoryAllocMB uint64  `json:"memory_alloc_mb"`
	MemoryTotalMB uint64  `json:"memory_total_mb"`
	CPUUsage      float64 `json:"cpu_usage_percent"`
	CPUCores      int     `json:"cpu_cores"`
	CPUUser       uint64  `json:"cpu_user"`
	CPUSystem     uint64  `json:"cpu_system"`
	CPUIdle       uint64  `json:"cpu_idle"`
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		uptime := time.Since(startTime).Round(time.Second).String()

		var CPUUsage float64
		var CPUUser, CPUSystem, CPUIdle uint64

		stat, err := linux.ReadStat("/proc/stat")
		if err != nil {
			log.Printf("Error reading CPU stats: %v", err)
		} else if len(stat.CPUStats) > 0 {
			cpuStat := stat.CPUStats[0]
			total := cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.Idle + cpuStat.IOWait + cpuStat.IRQ + cpuStat.SoftIRQ + cpuStat.Steal
			idle := cpuStat.Idle + cpuStat.IOWait

			if total > 0 {
				CPUUsage = float64(total-idle) / float64(total) * 100.0
			}

			CPUUser = cpuStat.User
			CPUSystem = cpuStat.System
			CPUIdle = cpuStat.Idle
		}

		response := HealthResponse{
			Time:          time.Now().Format(time.DateTime),
			Hostname:      hostname,
			PID:           os.Getpid(),
			Status:        "OK",
			GoVersion:     runtime.Version(),
			Uptime:        uptime,
			MemoryUsageMB: m.Alloc / 1024 / 1024,
			MemoryAllocMB: m.TotalAlloc / 1024 / 1024,
			MemoryTotalMB: m.Sys / 1024 / 1024,
			CPUUsage:      CPUUsage,
			CPUCores:      runtime.NumCPU(),
			CPUUser:       CPUUser,
			CPUSystem:     CPUSystem,
			CPUIdle:       CPUIdle,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	log.Print("Service UP » Listening on port 8443 !")
	if err := http.ListenAndServeTLS(":8443", "localhost.pem", "localhost-key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
