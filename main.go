package main

import (
	"fleetpulse/server"
	"fleetpulse/vehicle"
)

func main() {
	go vehicle.StartFleet(50)
	go server.MonitorAlerts()
	server.StartServer()
}
