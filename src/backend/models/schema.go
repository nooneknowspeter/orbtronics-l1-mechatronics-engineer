package models

import (
	"nooneknows/orbtronics-l1-mechatronics-engineer/helpers"
)

type DeviceData struct {
	Name        string  `json:"deviceName"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}

func SetupSchema() {
	schemaContent, _ := helpers.ReadSQLFile("db/migrations/schema.sql")

	DB.Exec(schemaContent)
}
