package utils

import (
	"fmt"
	"os"
	"path"
	"time"
)

const (
	DIR_PERMISSIONS = 0700
	CONFIG_EXT      = ".json"
)

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadWholeFile(filePath string, fileName string) ([]byte, error) {
	contents, err := os.ReadFile(path.Join(filePath, fileName))

	if err != nil {
		fmt.Printf("Unable to find file %s in path %s \n", fileName, filePath)
		return nil, err
	}

	return contents, nil
}

func CreateNewFile(filePath string, fileName string, fileContents string, channel chan error) {
	file, err := os.Create(path.Join(filePath, fileName))

	if err != nil {
		fmt.Printf("Unable to create the file %s in path %s", fileName, filePath)
		channel <- err
		return
	}

	defer file.Close()

	_, writeErr := file.WriteString(fileContents)

	if writeErr != nil {
		fmt.Printf("Unable to write to the file %s in path %s", fileName, filePath)
		channel <- writeErr
		return
	}

	channel <- nil
}

func CreateDirectory(dirPath string, dirName string) error {
	err := os.Mkdir(path.Join(dirPath, dirName), DIR_PERMISSIONS)

	return err
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
