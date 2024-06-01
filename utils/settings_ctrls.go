package utils

import (
	"encoding/json"
	"path"
)

const (
	GRUNT_SETTINGS = "grunt.json"
)

type Settings struct {
	Version string `json:"version"`
	Root    string `json:"root"`
	Configs string `json:"configs"`
	Content string `json:"content"`
	Logs    string `json:"logs"`
}

func CreateNewSettings(version string, root string, configs string, content string, logs string) *Settings {
	return &Settings{
		Version: version,
		Root:    root,
		Configs: configs,
		Content: content,
		Logs:    logs,
	}
}

func LoadGruntSettings(homeDir string) (*Settings, error) {
	settings, err := ReadWholeFile(path.Join(homeDir, ".grunt"), GRUNT_SETTINGS)

	if err != nil {
		return nil, err
	}

	var gruntSettings *Settings

	err = json.Unmarshal(settings, &gruntSettings)

	if err != nil {
		return nil, err
	}

	return gruntSettings, nil
}
