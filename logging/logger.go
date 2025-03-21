// Package logging provides a logger and related methods.
package logging

import (
	"io"
	"log"
	"os"
)

var logger *log.Logger

// GetLogger retrieves an instance of an application logger which outputs
// to a file if the debug flag is enabled
func GetLogger() (*log.Logger, error) {
	if logger != nil {
		return logger, nil
	}

	var writer io.Writer
	var err error

	writer = io.Discard // by default, suppress all logging output
	debugModeEnabled := os.Getenv("DEBUG") == "1"
	if debugModeEnabled {
		writer, err = os.Create("revive.log")
		if err != nil {
			return nil, err
		}
	}

	logger = log.New(writer, "", log.LstdFlags)

	if !debugModeEnabled {
		// Clear all flags to skip log output formatting step to increase
		// performance somewhat if we're not logging anything
		logger.SetFlags(0)
	}

	logger.Println("Logger initialized")

	return logger, nil
}
