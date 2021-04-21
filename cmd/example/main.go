package main

import "github.com/gojaguar/logger"

func main() {
	log := logger.NewLogger()

	log.Debug("Debug message")
	log.Info("Info message")
	log.Warn("Warn message")
	log.Error("Error message")
}
