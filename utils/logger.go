package utils

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func InitLogger(env string) *Logger {
	if env == "production" {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
	} else {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}
