package main

import (
	"os"
	"path"

	"github.com/JGugino/grunt/cmd"
	"github.com/JGugino/grunt/utils"
)

const (
	VERSION = "1.2.2"

	I_INIT    = "init"
	I_CREATE  = "create"
	I_LOGS    = "log"
	I_VERSION = "version"
	I_INFO    = "info"
)

func main() {
	if len(os.Args) <= 1 {
		utils.PrintError("Invalid usage", false, false)
		utils.PrintError("grunt {identifier} {args}", false, false)
		utils.PrintError("* If the path is not defined, the current working directory is used.", true, false)
		os.Exit(0)
	}

	//Gets the command identifier
	cmdIdentifier := os.Args[1]
	cmdArgs := os.Args[2:]

	//Determines the users home dir
	homeDir, err := os.UserHomeDir()

	utils.HandleError(err, true)

	//check for .grunt folder and create if it doesn't exist
	rootDirExist := utils.PathExists(path.Join(homeDir, ".grunt"))

	if !rootDirExist && cmdIdentifier != I_INIT {
		utils.PrintError("grunt hasn't been initialized, please run `grunt init`", false, false)
		os.Exit(0)
	}

	//Determines the users current working directory
	workingDir, err := os.Getwd()

	utils.HandleError(err, true)

	var gruntSettings *utils.Settings

	if cmdIdentifier != I_INIT {
		//Loads the grunt settings from the 'grunt.json' file located inside the root .grunt folder
		gruntSettings, err = utils.LoadGruntSettings(homeDir)

		utils.HandleError(err, true)
	}

	switch cmdIdentifier {
	case I_INIT:
		//Run the init command to create the root grunt folders
		initCmd := cmd.InitCmd{}
		err := initCmd.Execute(homeDir, rootDirExist, VERSION)

		utils.HandleError(err, true)
	case I_CREATE:
		//Run the create command to create a template config file and content folder with the specified name
		createCmd := cmd.CreateCmd{}
		err := createCmd.Execute(cmdArgs, gruntSettings.Configs, gruntSettings.Content)

		utils.HandleError(err, true)

	case I_LOGS:
		//Run the logs command to print out the general or error logs, or the available configs
		logsCmd := cmd.LogsCmd{}
		err := logsCmd.Execute(gruntSettings.Configs, gruntSettings.Logs, gruntSettings.Content, cmdArgs)

		utils.HandleError(err, true)

	case I_VERSION:
		//Run the version command which will print out the version of the current grunt install
		versionCmd := cmd.VersionCmd{}
		err := versionCmd.Execute(gruntSettings.Version)

		utils.HandleError(err, true)

	case I_INFO:
		//Run the info command to print out more info about grunt and it's settings
		infoCmd := cmd.InfoCmd{}
		err := infoCmd.Execute(gruntSettings)

		utils.HandleError(err, true)

	default:
		//If no command is specified run the config exection command
		configCmd := cmd.ConfigCmd{}
		err := configCmd.Execute(cmdIdentifier, gruntSettings.Configs, workingDir)

		utils.HandleError(err, true)
	}
}
