package utils

import (
	"encoding/json"
	"fmt"
)

func CreateInitDirectoriesIfDontExist(homeDir string, rootFolder string, exist bool) error {
	if !exist {
		//Create root dir
		err := CreateDirectory(homeDir, ".grunt")

		if err != nil {
			return err
		}

		//create config/content/logs folders
		err = CreateDirectory(rootFolder, "configs")
		if err != nil {
			return err
		}

		err = CreateDirectory(rootFolder, "content")
		if err != nil {
			return err
		}

		err = CreateDirectory(rootFolder, "logs")
		if err != nil {
			return err
		}

		PrintInfo(fmt.Sprintf("Created '.grunt' directory inside your home directory, %s", rootFolder), true)
		return nil
	}

	PrintInfo(fmt.Sprintf(".grunt directory already exists inside your home directory, %s", rootFolder), true)
	return nil
}

func CreateGruntSettings(version string, rootFolder string, configsFolder string, contentFolder string, logsFolder string) error {
	//Creates a new grunt settings file in the root of the .grunt directory
	newSettings := CreateNewSettings(version, rootFolder, configsFolder, contentFolder, logsFolder)

	settingsContent, err := json.Marshal(newSettings)

	if err != nil {
		return err
	}

	var channel chan error

	go CreateNewFile(rootFolder, "grunt.json", string(settingsContent), channel)

	err = <-channel

	if err != nil {
		return err
	}

	return nil
}
