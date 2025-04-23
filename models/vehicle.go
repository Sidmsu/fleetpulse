package models

type VehicleData struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Health    string  `json:"health"`
	Timestamp int64   `json:"timestamp"`
}
