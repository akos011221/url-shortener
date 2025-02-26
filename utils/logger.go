package utils

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func InitLogger(env string) *Logger {
	l := log.New(os.Stdout, "", log.LstdFlags)
	if env != "production" {
		l.SetFlags(log.LstdFlags | log.Lshortfile)
	}
	return &Logger{l}
}

func (l *Logger) Info(msg string) {
	l.Printf("[INFO] %s", msg)
}

func (l *Logger) Error(msg string) {
	l.Printf("[ERROR] %s", msg)
}
