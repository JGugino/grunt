package cmd

import (
	"errors"
	"fmt"

	"github.com/JGugino/grunt/utils"
)

type CreateCmd struct {
}

func (cmd *CreateCmd) Execute(args []string, configFolder string, contentFolder string) error {
	//If there is no name return an error
	if len(args) <= 0 {
		return errors.New("no-create-name")
	}

	//Create a variable for the config and content directory name
	configName := utils.AddConfigExt(args[0])
	dirName := args[0]

	utils.PrintInfo(fmt.Sprintf("Starting config creation for '%s'", configName), true)

	channel := make(chan error)

	//The template content for the generated config, replaces the %s with the specified name
	configContent := `
	{
		"id":"%s",
		"flags": [],
		"args": [],
		"directories": [
			{
				"name":"example",
				"subDirectories": [],
				"files": [
					{
						"name": "example.txt",
						"content": "This is example content for this example config"
					}
				]
			}
		],
		"commands": [
			{
				"command": "ls",
				"args": [
					"example/"
				]
			}
		]

	}
	`

	//Create the new config file inside the config folder with the specified name
	go utils.CreateNewFile(configFolder, configName, fmt.Sprintf(configContent, dirName), channel)

	err := <-channel

	if err != nil {
		return err
	}

	utils.PrintAction(fmt.Sprintf("Config file '%s' has been created", configName), true)

	//Create the new directory inside then content folder with the specified name
	err = utils.CreateDirectory(contentFolder, dirName)

	if err != nil {
		return err
	}

	utils.PrintAction(fmt.Sprintf("Content folder '%s' has been created", dirName), true)

	utils.PrintInfo(fmt.Sprintf("Config creation for '%s' has completed", configName), true)

	return nil
}
