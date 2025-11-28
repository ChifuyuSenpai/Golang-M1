package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
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

// Constantes pour les états des agents
const (
	StatusUP      AgentStatus = "UP"
	StatusWARNING AgentStatus = "WARNING"
	StatusDOWN    AgentStatus = "DOWN"
)

type AgentState struct {
	Metrics    AgentMetrics `json:"metrics"`
	Status     AgentStatus  `json:"status"`
	LastSeenAt time.Time    `json:"last_seen_at"`
}

// init map
type Server struct {
	mu     sync.RWMutex
	agents map[string]*AgentState
}

func NewServer() *Server {
	return &Server{
		agents: make(map[string]*AgentState),
	}
}

// http status : https://pkg.go.dev/net/http#pkg-constants

// POST /metrics
func (s *Server) handleMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405", http.StatusMethodNotAllowed)
		return
	}

	var metrics AgentMetrics
	if err := json.NewDecoder(r.Body).Decode(&metrics); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.agents[metrics.AgentID] = &AgentState{
		Metrics:    metrics,
		Status:     StatusUP,
		LastSeenAt: time.Now(),
	}
	s.mu.Unlock()

	log.Printf("Received from: %s", metrics.AgentID)
	w.WriteHeader(http.StatusOK)
}

// GET /agents
func (s *Server) handleListAgents(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.agents)
}

// update status
func (s *Server) updateStatuses() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		s.mu.Lock()
		now := time.Now()
		for _, state := range s.agents { // for id, state := range s.agents {
			elapsed := now.Sub(state.LastSeenAt)
			if elapsed > 30*time.Second {
				state.Status = StatusDOWN
			} else if elapsed > 15*time.Second {
				state.Status = StatusWARNING
			} else {
				state.Status = StatusUP
			}
		}
		s.mu.Unlock()
	}
}

// init server
func main() {
	server := NewServer()

	http.HandleFunc("/metrics", server.handleMetrics)
	http.HandleFunc("/agents", server.handleListAgents)

	go server.updateStatuses()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
