package cmd

import (
	"errors"
	"fmt"
	"path"
	"strings"

	"github.com/JGugino/grunt/utils"
)

type LogsCmd struct {
}

func (cmd *LogsCmd) Execute(configsFolder string, logsFolder string, contentFolder string, args []string) error {
	if len(args) <= 0 {
		return errors.New("invalid-log-args")
	}

	logType := args[0]

	//Logs the general log to the terminal
	if logType == "general" {
		logContent, err := utils.ReadWholeFile(logsFolder, "general.log")

		if err != nil {
			return err
		}

		splitContent := strings.Split(string(logContent), "\n")

		for i, v := range splitContent {
			utils.PrintInfo(fmt.Sprintf("[%d] %s\n", i+1, v), false)

		}

		return nil
	} else if logType == "error" {
		logContent, err := utils.ReadWholeFile(logsFolder, "errors.log")

		if err != nil {
			return err
		}

		splitContent := strings.Split(string(logContent), "\n")

		for i, v := range splitContent {
			utils.PrintError(fmt.Sprintf("[%d] %s\n", i+1, v), false, false)
		}
		return nil
	} else if logType == "configs" {
		files, err := utils.GetFilesInDirectory(configsFolder)

		if err != nil {
			return err
		}

		utils.PrintInfo(fmt.Sprintf("Config Files - %s", configsFolder), true)
		utils.PrintInfo("---------------------------------", true)
		utils.PrintInfo("Config Name | Has Content Folder", true)
		utils.PrintInfo("---------------------------------", true)
		for id, f := range files {
			fileName := strings.Split(f, ".")[0]

			exists := utils.PathExists(path.Join(contentFolder, fileName))

			contentExists := "no"

			if exists {
				contentExists = "yes"
			}

			utils.PrintInfo(fmt.Sprintf("%d) %s | %s ", id+1, f, contentExists), true)
		}
		return nil
	}

	return errors.New("invalid-log-type")
}
