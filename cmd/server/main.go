package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)
type Telemetry struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Battery   float64 `json:"battery"`
	Speed     float64 `json:"speed"`
}

func heartbeatHandler(w http.ResponseWriter, r *http.Request) {
	var telemetry Telemetry

	err := json.NewDecoder(r.Body).Decode(&telemetry)
	if err != nil {
		http.Error(w, "invalid telemetry", http.StatusBadRequest)
		return
	}

	fmt.Printf(
		"Vehicle=%s Lat=%.4f Lon=%.4f Battery=%.1f Speed=%.1f\n",
		telemetry.VehicleID,
		telemetry.Latitude,
		telemetry.Longitude,
		telemetry.Battery,
		telemetry.Speed,
	)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "heartbeat received")
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "FleetGuard server is running")
	})
	http.HandleFunc("/api/v1/vehicles/heartbeat", heartbeatHandler)

	fmt.Println("FleetGuard server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("server error:", err)
	}
}