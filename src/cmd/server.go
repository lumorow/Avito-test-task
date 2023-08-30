package main

import (
	"Avito-test-task/internal/app"
)

const configPath = "config/config.yaml"

// @title Avito-test-task
// @version 1.0
// description API Server for Avito-test-task

// @host localhost:8080
// @BasePath /

func main() {
	app.Run(configPath)
}
