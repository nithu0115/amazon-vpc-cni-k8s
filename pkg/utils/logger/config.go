package logger

import (
	"os"
)

const (
	defaultLogFilePath = "/host/var/log/aws-routed-eni/ipamd.log"
	defaultLogLevel    = "Debug"
	binaryName         = "L-IPamD"
)

// Configuration stores the config for the logger
type Configuration struct {
	BinaryName   string
	LogLevel     string
	LogLocation  string
}

func LoadLogConfig() *Configuration {
    return &Configuration{
		LogLevel: getLogLevel(),
		LogLocation: getLogFileLocation(),
		BinaryName: binaryName,
	}
}

// GetLogFileLocation returns the log file path
func getLogFileLocation() string {
	logFilePath := os.Getenv(envLogFilePath)
	if logFilePath == "" {
		logFilePath = defaultLogFilePath
	}
	return logFilePath
}

func getLogLevel() string {
	logLevel := os.Getenv(envLogLevel)
	switch logLevel {
	case "":
		logLevel = defaultLogLevel
		return logLevel
	default:
		return logLevel
	}
}
