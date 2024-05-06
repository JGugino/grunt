package utils

import (
	"fmt"
	"os"
)

const (
	DIR_PERMISSIONS = 0700
	CONFIG_EXT      = ".json"
)

func HandleError(err error, fatal bool) {
	if err != nil {
		//TODO: Output to log file
		fmt.Println(err)
		if fatal {
			os.Exit(0)
		}
	}
}

func AddConfigExt(fileName string) string {
	return fileName + CONFIG_EXT
}
