package api

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	classname string
}

func NewLogger(classname string) *Logger {
	return &Logger{classname: classname}
}

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

func (l *Logger) logWithTime() string {
	currentTime := time.Now()
	return currentTime.Format("01/02/2006 3:04:05PM")
}

func (l *Logger) log(level, message string) {
	finalMessage := fmt.Sprintf("[%s %s - %s]: %s", l.logWithTime(), l.classname, level, message)

	fmt.Println(finalMessage)

	l.logToFile(finalMessage)
}

func (l *Logger) Log(level, message string) {
	l.log(level, message)
}

func (l *Logger) Error(message string) {
	l.Log("ERROR", message)
}

func (l *Logger) Debug(message string) {
	l.Log("DEBUG", message)
}

func (l *Logger) Info(message string) {
	l.Log("INFO", message)
}

func (l *Logger) Warning(message string) {
	l.Log("WARNING", message)
}
