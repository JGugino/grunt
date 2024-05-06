package main

import (
	"fmt"
	"os"
	"path"

	"github.com/JGugino/grunt/utils"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Invalid format")
		fmt.Println("Usage: grunt {configId}")
		os.Exit(0)
	}

	configId := os.Args[1]

	homeDir, err := os.UserHomeDir()

	utils.HandleError(err, true)

	workingDir, err := os.Getwd()

	utils.HandleError(err, true)

	rootFolder := path.Join(homeDir, ".grunt")
	configsFolder := path.Join(rootFolder, "configs")
	logsFolder := path.Join(rootFolder, "logs")

	//check for .grunt folder and create if it doesn't exist
	rootDirExist := utils.PathExists(rootFolder)

	if !rootDirExist {
		//Create root dir
		os.Mkdir(rootFolder, utils.DIR_PERMISSIONS)

		//create config/logs folders
		os.Mkdir(configsFolder, utils.DIR_PERMISSIONS)
		os.Mkdir(logsFolder, utils.DIR_PERMISSIONS)

		fmt.Printf("created '.grunt' directory inside your home directory, %s \n", rootFolder)
	}

	//load selected config from the .grunt/configs folder
	config, err := utils.LoadConfig(configsFolder, configId)

	utils.HandleError(err, false)

	fmt.Println(workingDir)

	fmt.Println(config.Id)
}
