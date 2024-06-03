package cmd

import (
	"path"

	"github.com/JGugino/grunt/utils"
)

type InitCmd struct {
	RootDirExist bool
	HomeDir      string
	Version      string
}

func (cmd *InitCmd) Start() error {

	//program folder paths
	rootFolder := path.Join(cmd.HomeDir, ".grunt")
	configsFolder := path.Join(rootFolder, "configs")
	logsFolder := path.Join(rootFolder, "logs")
	contentFolder := path.Join(rootFolder, "content")

	//Creates the root directories required for grunt
	err := utils.CreateInitDirectoriesIfDontExist(cmd.HomeDir, rootFolder, cmd.RootDirExist)

	if err != nil {
		return err
	}

	err = utils.CreateGruntSettings(cmd.Version, rootFolder, configsFolder, contentFolder, logsFolder)

	if err != nil {
		return err
	}

	return nil
}
