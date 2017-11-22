package telldus_local

import (
	"github.com/prometheus/client_golang/prometheus"
)

// TellStickHost represents a Local Tellstick api
type TellStickHost struct {
	Address string
	Token   string
}

type Exporter struct {
	Metrics map[string]*prometheus.Desc
	TellStickHost
}

type SensorInfo struct {
	ID       int          `json:"id"`
	Model    string       `json:"model"`
	Name     string       `json:"name"`
	Novalues bool         `json:"novalues,omitempty"`
	Protocol string       `json:"protocol"`
	SensorID int          `json:"sensorId"`
	Battery  int          `json:"battery,omitempty"`
	Data     []SensorData `json:"data,omitempty"`
}

type SensorList struct {
	Sensor []SensorInfo `json:"sensor"`
}

type SensorData struct {
	Name  string  `json:"name"`
	Scale int     `json:"scale"`
	Value float64 `json:"value"`
}

type DeviceInfo struct {
	ID         int    `json:"id"`
	Methods    int    `json:"methods"`
	Name       string `json:"name"`
	State      int    `json:"state"`
	Statevalue string `json:"statevalue"`
	Type       string `json:"type"`
}

type DeviceList struct {
	Device []DeviceInfo `json:"device"`
}
