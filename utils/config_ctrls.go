package utils

import (
	"encoding/json"
)

type ConfigFile struct {
	Id          string      `json:"id"`
	Directories []Directory `json:"directories"`
	Commands    []Command   `json:"commands"`
}

type Directory struct {
	Name           string   `json:"name"`
	SubDirectories []string `json:"subDirectories"`
}

type Command struct {
	Command string   `json:"command"`
	Args    []string `json:"args"`
}

func LoadConfig(path string, configId string) (*ConfigFile, error) {
	configContents, err := ReadWholeFile(path, AddConfigExt(configId))

	if err != nil {
		return nil, err
	}

	var configFile *ConfigFile

	err = json.Unmarshal(configContents, &configFile)

	return configFile, err
}
