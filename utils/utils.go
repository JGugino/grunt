package utils

import (
	"fmt"
	"os"
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
