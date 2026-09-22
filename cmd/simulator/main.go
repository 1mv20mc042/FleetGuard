package main

import (
	"bytes"
	"net/http"
	"encoding/json"
	"fmt"
	"time"
)

type Vehicle struct {
	ID        string
	Latitude  float64
	Longitude float64
	Battery   float64
	Speed     float64
}

type Telemetry struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Battery   float64 `json:"battery"`
	Speed     float64 `json:"speed"`
}

func (v *Vehicle) Update() {
	v.Latitude += 0.0001
	v.Longitude += 0.0001
	v.Battery -= 0.1
	v.Speed = 12.5
}

func (v *Vehicle) GetTelemetry() Telemetry {
	return Telemetry{
		VehicleID: v.ID,
		Latitude:  v.Latitude,
		Longitude: v.Longitude,
		Battery:   v.Battery,
		Speed:     v.Speed,
	}
}

func main() {
	boat := Vehicle{
		ID:        "boat-001",
		Latitude:  19.0760,
		Longitude: 72.8777,
		Battery:   100,
		Speed:     0,
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		boat.Update()

		telemetry := boat.GetTelemetry()
		telemetryJSON, err:= json.Marshal(telemetry)
		if err != nil {
			fmt.Println("Error marshaling telemetry:", err)
			continue
		}

		resp, err := http.Post("http://localhost:8080/api/v1/vehicles/heartbeat", "application/json", bytes.NewBuffer(telemetryJSON))
		if err != nil {
			fmt.Println("Error sending telemetry:", err)
			continue
		}
		resp.Body.Close()
		fmt.Println("Telemetry sent")
		
	}
}