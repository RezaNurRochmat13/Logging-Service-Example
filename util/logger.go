package util

import "log"

func LoggerOutput(message string, info string, errorMessage string) {
	log.Printf("Logger info %s", info, "message : %s", message, "error : %s", errorMessage)
}
