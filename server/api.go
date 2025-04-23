package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"fleetpulse/models"
)

var (
	VehicleState = make(map[string]models.VehicleData)
	StateMutex   = sync.Mutex{}
)

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

func HandleState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	StateMutex.Lock()
	defer StateMutex.Unlock()
	json.NewEncoder(w).Encode(VehicleState)
}

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/telemetry", HandleTelemetry)
	http.HandleFunc("/state", HandleState)

	log.Println("FleetPulse server running on port", port)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
