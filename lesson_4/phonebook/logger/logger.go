package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	warnLogger = log.New(os.Stdout, "WARN: ", log.LstdFlags)
	infoLogger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
)

func Info(message string) {
	infoLogger.Println(message)
}

func Warn(err error, messages ...string) {
	if err == nil {
		return
	}

	message := "An error occurred"
	if len(messages) > 0 {
		message = fmt.Sprintf("%s: %s", message, strings.Join(messages, " "))
	}

	warnLogger.Printf("%s: %v", message, err)
}
