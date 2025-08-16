package main

import (
	"nooneknows/orbtronics-l1-mechatronics-engineer/helpers"
	"nooneknows/orbtronics-l1-mechatronics-engineer/models"
	"nooneknows/orbtronics-l1-mechatronics-engineer/routes"
)

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
