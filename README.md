# FleetPulse 🚗📡 (Go Backend)

FleetPulse is a real-time vehicle telemetry backend system built in Go. It simulates a fleet of 50 vehicles streaming GPS and health data to a centralized server. The server ingests data via REST APIs, logs alerts for critical/offline vehicles, and exposes a live state endpoint for visualization.

## 🌟 Features

- Real-time simulation of 50 vehicles using goroutines
- Vehicle telemetry includes GPS location and health status
- REST API to ingest `/telemetry` and expose `/state`
- Logs alerts for:
  - 🚨 CRITICAL health
  - ⚠️ OFFLINE vehicles (no data for >15s)
- CORS-enabled for frontend integration
- Easily extendable with persistence, metrics, or alerting systems

## 🛠 Tech Stack

- **Language**: Go (Golang)
- **Concurrency**: Goroutines
- **Transport**: REST over HTTP
- **Data Format**: JSON
- **Design**: In-memory storage for simplicity and speed

## 📦 Project Structure

fleetpulse/
├── main.go
├── go.mod
├── models/ # Vehicle struct
├── server/ # API endpoints + alert logic
└── vehicle/ # Simulated vehicle fleet


## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/fleetpulse.git
cd fleetpulse
```
### 2. Run the Go server
```bash
go run .

The server runs on http://localhost:8080
```
📡 API Endpoints
POST /telemetry
Receives JSON from simulated vehicles.

GET /state
Returns the current state of all vehicles as JSON.

💡 Example Output
```bash
{
  "vehicle-01": {
    "id": "vehicle-01",
    "latitude": 37.7754,
    "longitude": -122.4178,
    "health": "OK",
    "timestamp": 1713283502
  },
  ...
}





