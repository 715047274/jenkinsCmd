package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", log.LstdFlags)

func Info(message string) {
	logger.Println("[INFO]:" + message)
}

func Error(message string) {
	logger.Println("[ERROR]:" + message)
}
