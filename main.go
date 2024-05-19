package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/JGugino/grunt/cmd"
	"github.com/JGugino/grunt/utils"
)

const (
	I_INIT   = "init"
	I_CREATE = "create"
	I_LOGS   = "logs"
)

func main() {
	if len(os.Args) <= 1 {
		utils.PrintError("Invalid usage", false)
		utils.PrintError("grunt {identifier} {args}", false)
		utils.PrintError("* If the path is not defined, the current working directory is used.", true)
		os.Exit(0)
	}

	//Gets the command identifier
	cmdIdentifier := os.Args[1]

	//Determines the users home dir
	homeDir, err := os.UserHomeDir()

	utils.HandleError(err, true)

	//Determines the users current working directory
	workingDir, err := os.Getwd()

	utils.HandleError(err, true)

	//program folder paths
	rootFolder := path.Join(homeDir, ".grunt")
	configsFolder := path.Join(rootFolder, "configs")

	//check for .grunt folder and create if it doesn't exist
	rootDirExist := utils.PathExists(rootFolder)

	utils.HandleError(err, true)

	switch cmdIdentifier {
	case I_INIT:
		initCmd := cmd.InitCmd{}
		err := initCmd.Execute(homeDir, rootFolder, rootDirExist)

		utils.HandleError(err, true)
	case I_CREATE:
		createCmd := cmd.CreateCmd{}
		err := createCmd.Execute()

		utils.HandleError(err, true)

	case I_LOGS:
		logsCmd := cmd.LogsCmd{}
		err := logsCmd.Execute()

		utils.HandleError(err, true)

	default:
		//Get the time at the start of the config execution
		startingTime := time.Now()

		//load selected config from the .grunt/configs folder
		config, err := utils.LoadConfig(configsFolder, cmdIdentifier)

		if err != nil {
			utils.PrintError(fmt.Sprintf("Unable to find config '%s' inside of %s", cmdIdentifier, configsFolder), false)
			os.Exit(0)
		}

		utils.PrintInfo(fmt.Sprintf("Config '%s' has been loaded", cmdIdentifier))

		//execute config inside current working directory if a path isn't defined
		createPath, err := utils.GrabArgFromSlice(os.Args, "-p")

		//If there is no path argument defined assign the createPath to the current working directory
		if err != nil {
			createPath.Value = workingDir
		}

		//Parse all flags defined inside the config and assign them to an ActiveFlags struct for easy use
		flags := config.DetermineFlags()

		//Determine all of the required args inside the config and attempt to assign them to the provided values from the passed command
		var commandArgs []utils.CommandArg

		for _, arg := range config.Args {
			//Attempts to grab the defined argument from the os.Args slice
			cmd, err := utils.GrabArgFromSlice(os.Args, arg)

			//If it is not found it will display a warning in the terminal and log to the general log file
			if err != nil {
				utils.PrintWarning(fmt.Sprintf("Defined arg '%s' is unused", arg))
				return
			}

			//If it exists add it to the slice of existing arguments
			commandArgs = append(commandArgs, cmd)
		}

		utils.PrintInfo(fmt.Sprintf("Starting grunt in '%s'", createPath.Value))

		if !flags.SkipCreation {

			//Check if there is a flag to skip directory creation
			if !flags.SkipDirs {
				//create specified directories from config
				err = config.CreateDirectories(createPath.Value, commandArgs)

				utils.HandleError(err, false)

				utils.PrintAction("Directories have been created")
			} else {
				utils.PrintInfo("## Skipping directory creation ##")
			}

			//Check if there is a flag to skip file creation
			if !flags.SkipFiles {
				//create specified files from config
				err = config.CreateFiles(createPath.Value, commandArgs)

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
			err = config.ExecuteCommands(createPath.Value, commandArgs)

			utils.HandleError(err, false)

			utils.PrintAction("All commands have been executed")
		} else {
			utils.PrintInfo("## Skipping command execution ##")
		}

		timeTook := time.Since(startingTime)

		utils.PrintInfo(fmt.Sprintf("Config '%s' execution has completed: %s", cmdIdentifier, timeTook.String()))
	}
}
