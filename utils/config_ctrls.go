package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
)

const (
	SKIP_FILES_FLAG    string = "skipFiles"
	SKIP_DIRS_FLAG     string = "skipDirs"
	SKIP_CREATION_FLAG string = "skipCreation"
	SKIP_COMMANDS_FLAG string = "skipCommands"
)

type ConfigFile struct {
	Id          string      `json:"id"`
	Flags       []string    `json:"flags"`
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

type ActiveFlags struct {
	SkipFiles    bool
	SkipDirs     bool
	SkipCreation bool
	SkipCommands bool
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

func (config *ConfigFile) DetermineFlags() ActiveFlags {
	return ActiveFlags{
		SkipFiles:    StringExistsInSlice(config.Flags, SKIP_FILES_FLAG),
		SkipDirs:     StringExistsInSlice(config.Flags, SKIP_DIRS_FLAG),
		SkipCreation: StringExistsInSlice(config.Flags, SKIP_CREATION_FLAG),
		SkipCommands: StringExistsInSlice(config.Flags, SKIP_COMMANDS_FLAG),
	}
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
				err := CreateNewFile(path.Join(createPath, parentDir.Name, subDir.Name), file.Name, file.Content)

				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (config *ConfigFile) ExecuteCommands(executePath string) error {
	for c := 0; c < len(config.Commands); c++ {
		command := config.Commands[c]

		os.Chdir(executePath)
		cmd := exec.Command(command.Command, command.Args...)
		err := cmd.Run()

		if err != nil {
			return err
		}

		PrintAction(fmt.Sprintf("Command '%s %s' has been executed", command.Command, TurnSliceIntoString(command.Args)))
	}
	return nil
}
