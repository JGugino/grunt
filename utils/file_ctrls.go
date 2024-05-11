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

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func ReadWholeFile(filePath string, fileName string) ([]byte, error) {
	contents, err := os.ReadFile(path.Join(filePath, fileName))

	if err != nil {
		fmt.Printf("Unable to find config %s in path %s \n", fileName, filePath)
		return nil, err
	}

	return contents, nil
}

func CreateNewFile(filePath string, fileName string, fileContents string) error {
	file, err := os.Create(path.Join(filePath, fileName))

	if err != nil {
		fmt.Printf("Unable to create the file %s in path %s", fileName, filePath)
		return err
	}

	defer file.Close()

	_, writeErr := file.WriteString(fileContents)

	if writeErr != nil {
		fmt.Printf("Unable to write to the file %s in path %s", fileName, filePath)
		return writeErr
	}

	return nil
}

func CreateDirectory(dirPath string, dirName string) error {
	err := os.Mkdir(path.Join(dirPath, dirName), DIR_PERMISSIONS)

	return err
}
