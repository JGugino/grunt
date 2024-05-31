package main

import (
	"os"
	"path"

	"github.com/JGugino/grunt/cmd"
	"github.com/JGugino/grunt/utils"
)

const (
	I_INIT    = "init"
	I_CREATE  = "create"
	I_LOGS    = "log"
	I_CONFIGS = "config"
)

func main() {
	if len(os.Args) <= 1 {
		utils.PrintError("Invalid usage", false, false)
		utils.PrintError("grunt {identifier} {args}", false, false)
		utils.PrintError("* If the path is not defined, the current working directory is used.", true, false)
		os.Exit(1)
	}

	//Gets the command identifier
	cmdIdentifier := os.Args[1]
	cmdArgs := os.Args[2:]

	//Determines the users home dir
	homeDir, err := os.UserHomeDir()

	utils.HandleError(err, true)

	//check for .grunt folder and create if it doesn't exist
	rootDirExist := utils.PathExists(path.Join(homeDir, ".grunt"))

	if !rootDirExist {
		utils.PrintError("grunt hasn't been initialized, please run `grunt init`", false, false)
	}

	//Determines the users current working directory
	workingDir, err := os.Getwd()

	utils.HandleError(err, true)

	//program folder paths
	rootFolder := path.Join(homeDir, ".grunt")
	configsFolder := path.Join(rootFolder, "configs")
	logsFolder := path.Join(rootFolder, "logs")
	contentFolder := path.Join(rootFolder, "content")

	switch cmdIdentifier {
	case I_INIT:
		//Run the init command to create the root grunt folders
		initCmd := cmd.InitCmd{}
		err := initCmd.Execute(homeDir, rootFolder, rootDirExist)

		utils.HandleError(err, true)
	case I_CREATE:
		//Run the create command to create a template config file and content folder with the specified name
		createCmd := cmd.CreateCmd{}
		err := createCmd.Execute(cmdArgs, configsFolder, contentFolder)

		utils.HandleError(err, true)

	case I_LOGS:
		//Run the logs command to print out the general or error logs, or the available configs
		logsCmd := cmd.LogsCmd{}
		err := logsCmd.Execute(configsFolder, logsFolder, contentFolder, cmdArgs)

		utils.HandleError(err, true)

	default:
		//If no command is specified run the config exection command
		configCmd := cmd.ConfigCmd{}
		err := configCmd.Execute(cmdIdentifier, configsFolder, workingDir)

		utils.HandleError(err, true)
	}
}
