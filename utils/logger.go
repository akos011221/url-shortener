package utils

import (
	"log"
	"os"
)

func InitLogger(env string) {
	if env == "production" {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags)
	} else {
		log.SetOutput(os.Stdout)
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
}
