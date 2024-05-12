package main

import (
	"fmt"
	"os"
	"path"

	"github.com/JGugino/grunt/utils"
	"github.com/TwiN/go-color"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(color.Ize(color.Red, "Invalid usage"))
		fmt.Println(color.Ize(color.Red, "grunt {configId} {*path}"))
		fmt.Println(color.InBlackOverRed("NOTE: if the path is not defined, the current working directory is used."))
		os.Exit(0)
	}

	//Gets the
	configId := os.Args[1]

	homeDir, err := os.UserHomeDir()

	utils.HandleError(err, true)

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

		fmt.Printf(color.InCyan("created '.grunt' directory inside your home directory, %s \n"), rootFolder)
	}

	//load selected config from the .grunt/configs folder
	config, err := utils.LoadConfig(configsFolder, configId)

	utils.HandleError(err, true)

	fmt.Printf(color.InBlue("config '%s' has been loaded \n"), configId)

	//execute config inside current working directory if a path isn't defined
	var createPath string

	if len(os.Args) >= 3 {
		createPath = os.Args[2]
	} else {
		createPath = workingDir
	}

	fmt.Printf(color.InBlue("starting grunt in '%s' \n"), createPath)

	//create specified directories from config
	err = config.CreateDirectories(createPath)

	utils.HandleError(err, false)

	fmt.Println(color.InGreen("directories have been created"))

	//create specified files from config
	err = config.CreateFiles(createPath)

	utils.HandleError(err, false)

	fmt.Println(color.InGreen("files have been created"))
}
