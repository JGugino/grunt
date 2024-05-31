package cmd

import "github.com/JGugino/grunt/utils"

type InitCmd struct {
}

func (cmd *InitCmd) Execute(homeDir string, rootFolder string, rootDirExist bool) error {
	err := utils.CreateInitDirectoriesIfDontExist(homeDir, rootFolder, rootDirExist)

	if err != nil {
		return err
	}

	return nil
}
