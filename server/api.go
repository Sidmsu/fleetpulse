package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"fleetpulse/models"
)

var (
	VehicleState = make(map[string]models.VehicleData)
	StateMutex   = sync.Mutex{}
)

// HandleTelemetry handles incoming vehicle data
func HandleTelemetry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var data models.VehicleData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	StateMutex.Lock()
	VehicleState[data.ID] = data
	StateMutex.Unlock()

	fmt.Fprintln(w, "Telemetry received")
}

// HandleState returns the latest state of all vehicles
func HandleState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	StateMutex.Lock()
	defer StateMutex.Unlock()
	json.NewEncoder(w).Encode(VehicleState)
}

// MonitorAlerts logs alerts based on health and inactivity
func MonitorAlerts() {
	for {
		time.Sleep(5 * time.Second)
		now := time.Now().Unix()

		StateMutex.Lock()
		for id, data := range VehicleState {
			if data.Health == "CRITICAL" {
				log.Printf("🚨 CRITICAL: %s at (%.4f, %.4f)\n", id, data.Latitude, data.Longitude)
			}
			if now-data.Timestamp > 15 {
				log.Printf("⚠️ OFFLINE: %s (no data in %ds)\n", id, now-data.Timestamp)
			}
		}
		StateMutex.Unlock()
	}
}

// StartServer initializes the HTTP server
func StartServer() {
	http.HandleFunc("/telemetry", HandleTelemetry)
	http.HandleFunc("/state", HandleState)

	// Get dynamic port for Render, default to 8080 for local dev
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("FleetPulse server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
