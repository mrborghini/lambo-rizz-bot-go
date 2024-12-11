package api

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	classname string
}

// Start a new logger with the name of the file or class or function
func NewLogger(classname string) *Logger {
	return &Logger{classname: classname}
}

// Generate a rizzbot.log and add message
func (l *Logger) logToFile(message string) {
	filepath := "./rizzbot.log"
	tryFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tryFile.Close()

	if _, err := tryFile.WriteString(message + "\n"); err != nil {
		fmt.Println(err)
		return
	}
}

// Get the time current time in An American format
func (l *Logger) logWithTime() string {
	currentTime := time.Now()
	return currentTime.Format("01/02/2006 3:04:05PM")
}

// Print a message and add it to ./rizzbot.log file
func (l *Logger) log(level, message string) {
	finalMessage := fmt.Sprintf("[%s %s - %s]: %s", l.logWithTime(), l.classname, level, message)

	fmt.Println(finalMessage)

	if os.Getenv("WRITE_LOGS") == "true" {
		l.logToFile(finalMessage)
	}
}

// log something with a custom level and any message
func (l *Logger) Log(level, message string) {
	l.log(level, message)
}

// Log an error
func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

// Log anything that needs to be debugged
func (l *Logger) Debug(message string) {
	l.Log("DEBUG", message)
}

// Log any information
func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

// Log a warning
func (l *Logger) Warning(message string) {
	l.Log("WARNING", message)
}
