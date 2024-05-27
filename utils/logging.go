package utils

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/TwiN/go-color"
)

// Handles errors and has the ability to exit the program if fatal
func HandleError(err error, fatal bool) {
	if err != nil {
		if err.Error() == "no-create-name" {
			PrintWarning("You must provide a name for the config", true)
			os.Exit(1)
		} else if err.Error() == "invalid-log-type" {
			PrintError("Invalid log type (general or error)", false, true)
			os.Exit(1)
		} else if err.Error() == "invalid-log-args" {
			PrintError("Invalid log args", false, true)
			os.Exit(1)
		}

		PrintError(err.Error(), false, true)
		if fatal {
			os.Exit(1)
		}
	}
}

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
func PrintError(msg string, urgent bool, log bool) {
	if urgent {
		fmt.Println(color.InBlackOverRed(msg))
		if log {
			AppendToLogFile("error", " [ERROR] - "+msg)
		}
		return
	}

	fmt.Println(color.InRed(msg))
	AppendToLogFile("error", " [ERROR] - "+msg)
}

// Prints program info in blue and logs it to the logs folder under 'general.log'
func PrintInfo(msg string, log bool) {
	fmt.Println(color.InBlue(msg))
	if log {
		AppendToLogFile("general", " [INFO] - "+msg)
	}
}

// Prints program action in green and logs it to the logs folder under 'general.log'
func PrintAction(msg string, log bool) {
	fmt.Println(color.InGreen(msg))
	if log {
		AppendToLogFile("general", " [ACTION] - "+msg)
	}
}

// Prints program warning in black on a yellow background and logs it to the logs folder under 'general.log'
func PrintWarning(msg string, log bool) {
	fmt.Println(color.InBlackOverYellow(msg))
	if log {
		AppendToLogFile("general", " [WARNING] - "+msg)
	}
}
