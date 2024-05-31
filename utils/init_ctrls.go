package utils

import (
	"fmt"
	"os"
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
		os.Exit(0)
	}

	PrintInfo(fmt.Sprintf(".grunt directory already exists inside your home directory, %s", rootFolder), true)
	return nil
}
