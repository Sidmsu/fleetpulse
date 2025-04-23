package server

import (
	"encoding/json"
	"fleetpulse/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	VehicleState = make(map[string]models.VehicleData)
	StateMutex   = sync.Mutex{}
)

func HandleTelemetry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 👈 Allow frontend access

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

func HandleState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // 👈 Allow frontend access

	StateMutex.Lock()
	defer StateMutex.Unlock()
	json.NewEncoder(w).Encode(VehicleState)
}

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

// ...

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local dev
	}
	log.Println("FleetPulse server running on port", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
