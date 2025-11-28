package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type AgentMetrics struct {
	AgentID       string    `json:"agent_id"`
	Hostname      string    `json:"hostname"`
	OS            string    `json:"os"`
	Architecture  string    `json:"architecture"`
	CPUCores      int       `json:"cpu_cores"`
	Goroutines    int       `json:"goroutines"`
	MemoryUsageMB uint64    `json:"memory_usage_mb"`
	MemoryAllocMB uint64    `json:"memory_alloc_mb"`
	Uptime        string    `json:"uptime"`
	Timestamp     time.Time `json:"timestamp"`
}

type AgentStatus string

type AgentState struct {
	Metrics AgentMetrics `json:"metrics"`
	Status  AgentStatus  `json:"status"`
}

func listAgents() {
	resp, err := http.Get("http://localhost:8080/agents")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var agents map[string]*AgentState
	if err := json.NewDecoder(resp.Body).Decode(&agents); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n=== Agents Status ===")
	for id, state := range agents {
		fmt.Printf("\n[%s] %s\n", state.Status, id)
		fmt.Printf("  Hostname: %s\n", state.Metrics.Hostname)
		fmt.Printf("  OS: %s/%s\n", state.Metrics.OS, state.Metrics.Architecture)
		fmt.Printf("  CPUs: %d\n", state.Metrics.CPUCores)
		fmt.Printf("  Memory: %d MB\n", state.Metrics.MemoryUsageMB)
		fmt.Printf("  Uptime: %s\n", state.Metrics.Uptime)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: client list")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		listAgents()
	default:
		fmt.Println("Unknown command test debug")
	}
}
