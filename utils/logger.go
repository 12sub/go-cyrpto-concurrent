package utils

// time to implement logging/verbose to my code

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	logger    *log.Logger
	logToFile bool
	logFile   *os.File
	logWriter io.Writer
	LogLevel	string
	isTerminal	bool
)

// ANSI color code
var (
	Reset	= 	"\033[0m"
	Red	= 	"\033[31m"
	Yellow	= 	"\033[33m"
	Cyan	= 	"\033[36m"
	Green	= 	"\033[32m"
	Gray	= 	"\033[90m"
)

// First step: initializing the logger
func InitLogger(level string, enableFile bool) {
	LogLevel = level
	logToFile = enableFile
	isTerminal = true


	// writing to log
	logWriter = os.Stdout
	// error handling
	if logToFile {
		var err error
		logFile, err := os.OpenFile("crypto-cli.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("[ERROR] Failed to create log file:", err)
			return
		} else {
			logWriter = io.MultiWriter(os.Stdout, logFile) // includes timestamps
		}
		// logger.SetOutput(logFile)
	}
	// creating new log output
	logger = log.New(logWriter, "", log.LstdFlags)
}

func shouldLog(req string) bool {
	order := map[string]int{"debug": 1, "info": 2, "warn": 3, "error": 4}
	return order[req] >= order[LogLevel]
}

// Function to create debug to debug logs
func Debug(msg string, args ...any) {
	if shouldLog("debug") {
		logWithColor("[DEBUG] ", Gray, msg, args...)
	}
}


// Function to create info to print regular logs
func Info(msg string, args ...any) {
	if shouldLog("info") {
		logWithColor("[INFO] ", Cyan, msg, args...)
	}
}

// Function to create info to print regular logs
func Warn(msg string, args ...any) {
	if shouldLog("warn") {
		logWithColor("[WARN] ", Yellow, msg, args...)
	}
}



// Function to create error to print error logs
func Error(msg string, args ...any) {
	if shouldLog("Error"){
		logWithColor("[ERROR] ", Red, msg, args...)
	}
}
func logWithColor(prefix, color, msg string, args ...any) {
	colorPrefix := prefix
	if isTerminal {
		colorPrefix = fmt.Sprintf("%s%s%s", color, prefix, Reset)
	}
	logger.Printf("%s %s", colorPrefix, fmt.Sprintf(msg, args...))
}
// Function to cleanup log file if opened
func Cleanup() {
	if logFile != nil {
		logFile.Close()
	}
}
