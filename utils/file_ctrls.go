package utils

import (
	"fmt"
	"os"
	"path"
)

const (
	DIR_PERMISSIONS = 0700
	CONFIG_EXT      = ".json"
)

// Returns true if the path exists and false if it doesnt
func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadWholeFile(filePath string, fileName string) ([]byte, error) {
	contents, err := os.ReadFile(path.Join(filePath, fileName))

	if err != nil {
		return nil, err
	}

	return contents, nil
}

func CreateNewFile(filePath string, fileName string, fileContents string, channel chan error) {
	file, err := os.Create(path.Join(filePath, fileName))

	if err != nil {
		PrintError(fmt.Sprintf("Unable to create the file %s in path %s", fileName, filePath), false, true)
		channel <- err
		return
	}

	defer file.Close()

	_, writeErr := file.WriteString(fileContents)

	if writeErr != nil {
		PrintError(fmt.Sprintf("Unable to write to the file %s in path %s", fileName, filePath), false, true)
		channel <- writeErr
		return
	}

	channel <- nil
}

func CreateDirectory(dirPath string, dirName string) error {
	err := os.Mkdir(path.Join(dirPath, dirName), DIR_PERMISSIONS)

	return err
}
