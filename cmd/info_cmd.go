package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/JGugino/grunt/utils"
)

type InfoCmd struct {
}

func (cmd *InfoCmd) Execute(gruntSetting *utils.Settings) error {
	exe, err := os.Executable()

	if err != nil {
		return err
	}

	exePath := filepath.Dir(exe)

	utils.PrintInfo(fmt.Sprintf("grunt - version %s", gruntSetting.Version), true)
	utils.PrintInfo(fmt.Sprintf("grunt location: %s", exePath), true)
	utils.PrintInfo("--------------------------", true)
	utils.PrintInfo(fmt.Sprintf("root: %s", gruntSetting.Root), true)
	utils.PrintInfo(fmt.Sprintf("configs: %s", gruntSetting.Configs), true)
	utils.PrintInfo(fmt.Sprintf("content: %s", gruntSetting.Content), true)
	utils.PrintInfo(fmt.Sprintf("logs: %s", gruntSetting.Logs), true)
	utils.PrintInfo("--------------------------", true)
	utils.PrintInfo("** settings can be modified inside of grunt.json **", true)
	return nil
}
