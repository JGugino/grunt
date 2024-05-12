package utils

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

func HandleError(err error, fatal bool) {
	if err != nil {
		//TODO: Output to log file
		PrintError(err.Error(), false)
		if fatal {
			os.Exit(0)
		}
	}
}

func AddConfigExt(fileName string) string {
	return fileName + CONFIG_EXT
}

func StringExistsInSlice(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}

	return false
}

func PrintError(msg string, urgent bool) {
	if urgent {
		fmt.Println(color.InBlackOverRed(msg))
		return
	}

	fmt.Println(color.InRed(msg))
}
func PrintInfo(msg string) {
	fmt.Println(color.InBlue(msg))
}
func PrintAction(msg string) {
	fmt.Println(color.InGreen(msg))
}
