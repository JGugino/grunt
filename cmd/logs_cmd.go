package cmd

import "fmt"

type LogsCmd struct {
}

func (cmd *LogsCmd) Execute() error {
	fmt.Println("Execute Logs Cmd")

	return nil
}
