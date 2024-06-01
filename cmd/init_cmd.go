package cmd

import (
	"path"

	"github.com/JGugino/grunt/utils"
)

type InitCmd struct {
}

func (cmd *InitCmd) Execute(homeDir string, rootDirExist bool, version string) error {

	//program folder paths
	rootFolder := path.Join(homeDir, ".grunt")
	configsFolder := path.Join(rootFolder, "configs")
	logsFolder := path.Join(rootFolder, "logs")
	contentFolder := path.Join(rootFolder, "content")

	//Creates the root directories required for grunt
	err := utils.CreateInitDirectoriesIfDontExist(homeDir, rootFolder, rootDirExist)

	if err != nil {
		return err
	}

	err = utils.CreateGruntSettings(version, rootFolder, configsFolder, contentFolder, logsFolder)

	if err != nil {
		return err
	}

	return nil
}
