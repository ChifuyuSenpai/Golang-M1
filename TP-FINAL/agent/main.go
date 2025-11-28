package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/c9s/goprocinfo/linux"
)

type AgentMetrics struct { // Reprise du struct HealthResponse du TP-HTTP-API
	ID            string    `json:"agent_id"`
	Hostname      string    `json:"hostname"`
	OS            string    `json:"os"`
	Architecture  string    `json:"architecture"`
	CPUCores      int       `json:"cpu_cores"`
	CPUUsage      float64   `json:"cpu_usage_percent"`
	CPUUser       uint64    `json:"cpu_user"`
	CPUSystem     uint64    `json:"cpu_system"`
	CPUIdle       uint64    `json:"cpu_idle"`
	MemoryUsageMB uint64    `json:"memory_usage_mb"`
	MemoryAllocMB uint64    `json:"memory_alloc_mb"`
	Uptime        string    `json:"uptime"`
	Timestamp     time.Time `json:"timestamp"`
}

var (
	startTime    = time.Now()
	serverURL    = "http://localhost:8080/metrics"
	sendInterval = 5 * time.Second
	ID           string
)

func collectMetrics() AgentMetrics {
	hostname, _ := os.Hostname()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

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

	return AgentMetrics{
		ID:            ID,
		Hostname:      hostname,
		OS:            runtime.GOOS,
		Architecture:  runtime.GOARCH,
		CPUCores:      runtime.NumCPU(),
		CPUUsage:      CPUUsage,
		CPUUser:       CPUUser,
		CPUSystem:     CPUSystem,
		CPUIdle:       CPUIdle,
		MemoryUsageMB: m.Alloc / 1024 / 1024,
		MemoryAllocMB: m.TotalAlloc / 1024 / 1024,
		Uptime:        time.Since(startTime).Round(time.Second).String(),
		Timestamp:     time.Now(),
	}
}

func sendMetrics() error {
	metrics := collectMetrics()
	data, err := json.Marshal(metrics)
	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(serverURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	log.Printf("Metrics sent (Status: %d)", resp.StatusCode)
	return nil
}

func main() {
	hostname, _ := os.Hostname()
	ID = hostname + "-" + time.Now().Format("20060102150405")

	log.Printf("Agent started (ID: %s)", ID)
	log.Printf("Sending to %s every %v", serverURL, sendInterval)

	ticker := time.NewTicker(sendInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := sendMetrics(); err != nil {
			log.Printf("Error: %v", err)
		}
	}
}
