package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type SystemStatus struct {
	App       string `json:"app"`
	Category  string `json:"category"`
	Tech      string `json:"tech"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		status := SystemStatus{
			App:       "smart-grid-energy-go-fiber-v2026-75",
			Category:  "Smart Grid Energy & Load Balancing",
			Tech:      "Go / Fiber Microservice",
			Timestamp: time.Now().Unix(),
		}
		json.NewEncoder(w).Encode(status)
	})

	fmt.Println("[smart-grid-energy-go-fiber-v2026-75] Service starting on port :8080...")
	http.ListenAndServe(":8080", nil)
}
