package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"src/global"
	"time"
)

// InitializeLogger sets up the logging to a file with the current date and time.
func InitializeLogger() (*os.File, error) {
	// Ensure the logs directory exists
	if err := os.MkdirAll(global.LogDir, os.ModePerm); err != nil {
		return nil, err
	}

	// Create log file with the current date and time
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	logFileName := filepath.Join(global.LogDir, currentTime+".txt")
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	// Set log output to both the file and the console
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	return logFile, nil
}

// CloseLogger closes the log file.
func CloseLogger(logFile *os.File) {
	if logFile != nil {
		logFile.Close()
	}
}
