package main

import (
	"nooneknows/orbtronics-l1-mechatronics-engineer/helpers"
	"nooneknows/orbtronics-l1-mechatronics-engineer/models"
	"nooneknows/orbtronics-l1-mechatronics-engineer/routes"
)

type DeviceData struct {
	Name        string  `json:"deviceName"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}
func main() {
	parentEnvFile := "../../.env"
	if err := helpers.LoadEnvironmentFile(parentEnvFile); err == nil {
	}

	cwdEnvFile := ".env"
	if err := helpers.LoadEnvironmentFile(cwdEnvFile); err == nil {
	}

	models.StartDatabaseConnection()

	routes.StartServer()
}
