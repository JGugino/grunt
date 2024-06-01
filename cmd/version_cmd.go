package cmd

import (
	"fmt"

	"github.com/JGugino/grunt/utils"
)

type VersionCmd struct {
}

func (cmd *VersionCmd) Execute(version string) error {
	utils.PrintInfo(fmt.Sprintf("grunt - version %s", version), true)
	return nil
}
