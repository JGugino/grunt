package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	SKIP_FILES_FLAG    string = "skipFiles"
	SKIP_DIRS_FLAG     string = "skipDirs"
	SKIP_CREATION_FLAG string = "skipCreation"
	SKIP_COMMANDS_FLAG string = "skipCommands"
)

type ConfigFile struct {
	Id          string          `json:"id"`
	Args        []string        `json:"args"`
	Flags       []string        `json:"flags"`
	Directories []Directory     `json:"directories"`
	Commands    []ConfigCommand `json:"commands"`
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

type ConfigCommand struct {
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

func DetermineFileContent(config ConfigFile, content string) (string, error) {
	hasPath, extracted := ExtractPathOrReturnOriginalContent(content)

	if hasPath {
		splitPath := strings.Split(extracted, string(os.PathSeparator))

		homedir, err := os.UserHomeDir()

		if err != nil {
			return "", err
		}

		filePath := path.Join(homedir, ".grunt", "content", config.Id)
		fileName := splitPath[len(splitPath)-1:][0]

		fileContent, err := ReadWholeFile(filePath, fileName)

		if err != nil {
			return "", err
		}

		return string(fileContent), nil
	}

	return content, nil
}

func (config *ConfigFile) DetermineFlags() *ActiveFlags {
	return &ActiveFlags{
		SkipFiles:    StringExistsInSlice(config.Flags, SKIP_FILES_FLAG),
		SkipDirs:     StringExistsInSlice(config.Flags, SKIP_DIRS_FLAG),
		SkipCreation: StringExistsInSlice(config.Flags, SKIP_CREATION_FLAG),
		SkipCommands: StringExistsInSlice(config.Flags, SKIP_COMMANDS_FLAG),
	}
}

func (config *ConfigFile) CreateDirectories(createPath string, definedArgs []CommandArg) error {
	for p := 0; p < len(config.Directories); p++ {
		parentDir := config.Directories[p]

		parentDir.Name = ReplaceArgWithValueInString(parentDir.Name, definedArgs)

		//Create parent directory
		err := CreateDirectory(createPath, parentDir.Name)

		if err != nil {
			return err
		}

		//Create sub directories
		for s := 0; s < len(parentDir.SubDirectories); s++ {
			subDir := parentDir.SubDirectories[s]
			err := CreateDirectory(path.Join(createPath, parentDir.Name), ReplaceArgWithValueInString(subDir.Name, definedArgs))

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (config *ConfigFile) CreateFiles(createPath string, definedArgs []CommandArg) error {
	for p := 0; p < len(config.Directories); p++ {
		parentDir := config.Directories[p]

		parentDir.Name = ReplaceArgWithValueInString(parentDir.Name, definedArgs)

		//Create parent directory files
		for pf := 0; pf < len(parentDir.Files); pf++ {
			file := parentDir.Files[pf]

			var extractedContent string

			extractedContent, err := DetermineFileContent(*config, file.Content)

			if err != nil {
				return err
			}

			channel := make(chan error)

			go CreateNewFile(path.Join(createPath, parentDir.Name), ReplaceArgWithValueInString(file.Name, definedArgs), extractedContent, channel)

			err = <-channel

			if err != nil {
				return err
			}
		}

		//Create sub directories files
		for s := 0; s < len(parentDir.SubDirectories); s++ {
			subDir := parentDir.SubDirectories[s]

			//Create sub directory files
			for sf := 0; sf < len(subDir.Files); sf++ {
				file := subDir.Files[sf]

				var extractedContent string

				extractedContent, err := DetermineFileContent(*config, file.Content)

				if err != nil {
					return err
				}

				channel := make(chan error)

				go CreateNewFile(path.Join(createPath, parentDir.Name, ReplaceArgWithValueInString(subDir.Name, definedArgs)), ReplaceArgWithValueInString(file.Name, definedArgs), extractedContent, channel)

				err = <-channel

				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (config *ConfigFile) ExecuteCommands(executePath string, definedArgs []CommandArg) error {
	os.Chdir(executePath)

	for c := 0; c < len(config.Commands); c++ {
		command := config.Commands[c]

		cmdArgs := ReplaceArgWithValueInSlice(command.Args, definedArgs)

		channel := make(chan CommandReturn)

		go ExecuteCommand(channel, command.Command, cmdArgs)

		commandReturn := <-channel

		if commandReturn.Err != nil {
			PrintError(fmt.Sprintf("Command '%s %s' has failed to execute", command.Command, TurnSliceIntoString(cmdArgs)), false, true)
			return commandReturn.Err
		}

		PrintAction(fmt.Sprintf("Command '%s %s' has been executed", command.Command, TurnSliceIntoString(cmdArgs)), true)
		if len(commandReturn.Output) > 0 {
			PrintAction(fmt.Sprintf("\n###OUTPUT###\n%s", commandReturn.Output), true)
		}
	}
	return nil
}
