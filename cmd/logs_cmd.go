package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JGugino/grunt/utils"
)

type LogsCmd struct {
}

func (cmd *LogsCmd) Execute(logsFolder string, args []string) error {
	if len(args) <= 0 {
		return errors.New("invalid-log-args")
	}

	logType := args[0]

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
	}

	return errors.New("invalid-log-type")
}
