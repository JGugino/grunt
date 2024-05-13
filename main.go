package main

import (
	"fmt"
	"os"
	"path"

	"github.com/JGugino/grunt/utils"
)

func main() {
	if len(os.Args) <= 1 {
		utils.PrintError("Invalid usage", false)
		utils.PrintError("grunt {configId} {path*}", false)
		utils.PrintError("* If the path is not defined, the current working directory is used.", true)
		os.Exit(0)
	}

	//Gets the config id to load
	configId := os.Args[1]

	//Determines the users home dir
	homeDir, err := os.UserHomeDir()

	utils.HandleError(err, true)

	//Determines the users current working directory
	workingDir, err := os.Getwd()

	utils.HandleError(err, true)

	//program folder paths
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

		utils.PrintInfo(fmt.Sprintf("Created '.grunt' directory inside your home directory, %s", rootFolder))
	}

	//load selected config from the .grunt/configs folder
	config, err := utils.LoadConfig(configsFolder, configId)

	utils.HandleError(err, true)

	utils.PrintInfo(fmt.Sprintf("Config '%s' has been loaded", configId))

	//execute config inside current working directory if a path isn't defined
	var createPath string

	if len(os.Args) >= 3 {
		createPath = os.Args[2]
	} else {
		createPath = workingDir
	}

	flags := config.DetermineFlags()

	utils.PrintInfo(fmt.Sprintf("Starting grunt in '%s'", createPath))

	if !flags.SkipCreation {

		//Check if there is a flag to skip directory creation
		if !flags.SkipDirs {
			//create specified directories from config
			err = config.CreateDirectories(createPath)

			utils.HandleError(err, false)

			utils.PrintAction("Directories have been created")
		} else {
			utils.PrintInfo("## Skipping directory creation ##")
		}

		//Check if there is a flag to skip file creation
		if !flags.SkipFiles {
			//create specified files from config
			err = config.CreateFiles(createPath)

			utils.HandleError(err, false)

			utils.PrintAction("Files have been created")
		} else {
			utils.PrintInfo("## Skipping file creation ##")
		}
	} else {
		utils.PrintInfo("## Skipping directory & file creation ##")
	}

	//Checks if there is a flag to skip command execution
	if !flags.SkipCommands {
		err = config.ExecuteCommands(createPath)

		utils.HandleError(err, false)

		utils.PrintAction("All commands have been executed")
	} else {
		utils.PrintInfo("## Skipping command execution ##")
	}

}
