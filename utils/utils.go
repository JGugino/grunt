package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

func HandleError(err error, fatal bool) {
	if err != nil {
		PrintError(err.Error(), false)
		if fatal {
			os.Exit(0)
		}
	}
}

func AddConfigExt(fileName string) string {
	return fileName + CONFIG_EXT
}

func TurnSliceIntoString(slice []string) string {
	var output string

	for _, v := range slice {
		output += v + " "
	}

	return strings.Trim(output, " ")
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
		AppendToLogFile("error", msg)
		return
	}

	fmt.Println(color.InRed(msg))
	AppendToLogFile("error", msg)
}
func PrintInfo(msg string) {
	fmt.Println(color.InBlue(msg))
	AppendToLogFile("general", msg)
}
func PrintAction(msg string) {
	fmt.Println(color.InGreen(msg))
	AppendToLogFile("general", msg)
}
