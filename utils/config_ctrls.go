package utils

import (
	"encoding/json"
	"path"
)

type ConfigFile struct {
	Id          string      `json:"id"`
	Directories []Directory `json:"directories"`
	Commands    []Command   `json:"commands"`
}

type Directory struct {
	Name           string         `json:"name"`
	SubDirectories []SubDirectory `json:"subDirectories"`
	Files          []Files        `json:"files"`
}

type SubDirectory struct {
	Name  string  `json:"name"`
	Files []Files `json:"files"`
}

type Files struct {
	Name    string `json:"name"`
	Content string `json:"content"`
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

func (config *ConfigFile) CreateDirectories(createPath string) error {
	for p := 0; p < len(config.Directories); p++ {
		parentDir := config.Directories[p]
		//Create parent directory
		err := CreateDirectory(createPath, parentDir.Name)

		if err != nil {
			return err
		}

		//Create sub directories
		for s := 0; s < len(parentDir.SubDirectories); s++ {
			subDir := parentDir.SubDirectories[s]
			err := CreateDirectory(path.Join(createPath, parentDir.Name), subDir.Name)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (config *ConfigFile) CreateFiles(createPath string) error {
	for p := 0; p < len(config.Directories); p++ {
		parentDir := config.Directories[p]

		//Create parent directory files
		for pf := 0; pf < len(parentDir.Files); pf++ {
			file := parentDir.Files[pf]
			err := CreateNewFile(path.Join(createPath, parentDir.Name), file.Name, file.Content)

			if err != nil {
				return err
			}
		}

		//Create sub directories files
		for s := 0; s < len(parentDir.SubDirectories); s++ {
			subDir := parentDir.SubDirectories[s]

			//Create sub directory files
			for pf := 0; pf < len(parentDir.Files); pf++ {
				file := parentDir.Files[pf]
				err := CreateNewFile(path.Join(path.Join(createPath, parentDir.Name), subDir.Name), file.Name, file.Content)

				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
