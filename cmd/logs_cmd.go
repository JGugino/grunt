package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/JGugino/grunt/utils"
	"github.com/TwiN/go-color"
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
			fmt.Printf(color.InBlue("[%d] %s\n"), i+1, v)
		}

		return nil
	} else if logType == "error" {
		logContent, err := utils.ReadWholeFile(logsFolder, "errors.log")

		if err != nil {
			return err
		}

		splitContent := strings.Split(string(logContent), "\n")

		for i, v := range splitContent {
			fmt.Printf(color.InRed("[%d] %s\n"), i+1, v)
		}
		return nil
	}

	return errors.New("invalid-log-type")
}
