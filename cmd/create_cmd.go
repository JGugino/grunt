package cmd

import "fmt"

type CreateCmd struct {
}

func (cmd *CreateCmd) Execute() error {
	fmt.Println("Execute Create Cmd")

	return nil
}
