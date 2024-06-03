package utils

import (
	"errors"
	"fmt"
)

type CommandHandler struct {
	RootExist  bool
	Identifier string
	Flags      []string
	Commands   map[string]Command
}

type Command struct {
	Identifier string
	Flags      []string
	Execute    interface{ Start() error }
}

func CreateNewCommandHandler(identifier string, flags []string, commands map[string]Command) *CommandHandler {
	return &CommandHandler{
		Identifier: identifier,
		Flags:      flags,
		Commands:   commands,
	}
}

func (cmd *CommandHandler) AddCmd(command Command) error {
	if _, ok := cmd.Commands[command.Identifier]; ok {
		return errors.New("command-exists")
	}

	cmd.Commands[command.Identifier] = command

	return nil
}

func (cmd *CommandHandler) ExecuteCommand() error {
	if !cmd.RootExist {
		return errors.New("no-root")
	}

	command, ok := cmd.Commands[cmd.Identifier]

	if ok {
		err := command.Execute.Start()
		if err != nil {
			return err
		}
	} else {
		fmt.Print("Run Config")
		return nil
	}
	return nil
}
