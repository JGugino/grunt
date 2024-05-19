package cmd

import "github.com/JGugino/grunt/utils"

type InitCmd struct {
}

func (cmd *InitCmd) Execute(homeDir string, rootFolder string, rootDirExist bool) error {
	return utils.CreateInitDirectoriesIfDontExist(homeDir, rootFolder, rootDirExist)
}
