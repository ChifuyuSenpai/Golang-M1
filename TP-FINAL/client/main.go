package main

import (
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