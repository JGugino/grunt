package utils

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/TwiN/go-color"
)

func AppendToLogFile(logType string, logContent string) error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	errorLogPath := path.Join(homeDir, ".grunt", "logs", "errors.log")
	generalLogPath := path.Join(homeDir, ".grunt", "logs", "general.log")

	var selectedLogPath string

	if logType == "error" {
		selectedLogPath = errorLogPath
	} else if logType == "general" {
		selectedLogPath = generalLogPath
	}

	log, err := os.OpenFile(selectedLogPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer log.Close()

	timestampedContent := time.Now().Format("Mon Jan _2 15:04:05 2006") + " - " + logContent + "\n"

	if _, err := log.WriteString(timestampedContent); err != nil {
		return err
	}

	return nil
}

// Prints an error in black over a red background and logs it to the logs folder under 'errors.log'
func PrintErrorAndLog(msg string, urgent bool) {
	if urgent {
		fmt.Println(color.InBlackOverRed(msg))
		AppendToLogFile("error", " [ERROR] - "+msg)
		return
	}

	fmt.Println(color.InRed(msg))
	AppendToLogFile("error", " [ERROR] - "+msg)
}

// Prints program info in blue and logs it to the logs folder under 'general.log'
func PrintInfoAndLog(msg string) {
	fmt.Println(color.InBlue(msg))
	AppendToLogFile("general", " [INFO] - "+msg)
}

// Prints program action in green and logs it to the logs folder under 'general.log'
func PrintActionAndLog(msg string) {
	fmt.Println(color.InGreen(msg))
	AppendToLogFile("general", " [ACTION] - "+msg)
}

// Prints program warning in black on a yellow background and logs it to the logs folder under 'general.log'
func PrintWarningAndLog(msg string) {
	fmt.Println(color.InBlackOverYellow(msg))
	AppendToLogFile("general", " [WARNING] - "+msg)
}
