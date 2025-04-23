package vehicle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"fleetpulse/models"
)

func randomCoord(base float64) float64 {
	return base + (rand.Float64()-0.5)/100
}

func getHealth() string {
	status := []string{"OK", "WARN", "CRITICAL"}
	return status[rand.Intn(len(status))]
}

func SimulateVehicle(id string) {
	for {
		data := models.VehicleData{
			ID:        id,
			Latitude:  randomCoord(37.7749),
			Longitude: randomCoord(-122.4194),
			Health:    getHealth(),
			Timestamp: time.Now().Unix(),
		}

		jsonData, _ := json.Marshal(data)
		_, err := http.Post("http://localhost:8080/telemetry", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("❌ Error sending telemetry for", id)
		}

		time.Sleep(2 * time.Second)
	}
}

func StartFleet(n int) {
	for i := 1; i <= n; i++ {
		go SimulateVehicle(fmt.Sprintf("vehicle-%02d", i))
	}
}
