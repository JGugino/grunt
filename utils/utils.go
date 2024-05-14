package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/TwiN/go-color"
)

type CommandArg struct {
	Name  string
	Value string
}

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

func GrabArgFromSlice(slice []string, arg string) (CommandArg, error) {
	for _, v := range slice {
		splitArg := strings.Split(v, "=")
		if len(splitArg) > 1 {
			if splitArg[0] == arg {
				return CommandArg{Name: splitArg[0], Value: splitArg[1]}, nil
			}
		}
	}

	return CommandArg{}, errors.New("not-found")
}

func ReplaceArgWithValueInSlice(slice []string, args []CommandArg) []string {
	replacedSlice := make([]string, 0)
	for _, s := range slice {
		for _, a := range args {
			argStart := strings.Index(s, "{")
			argEnd := strings.Index(s, "}")

			var extractedArg string

			if argStart != -1 && argEnd != -1 {
				extractedArg = s[argStart:argEnd] + "}"
			} else {
				extractedArg = s
			}

			if extractedArg == "{"+a.Name+"}" {
				replacedSlice = append(replacedSlice, s[:argStart]+a.Value+s[argEnd+1:])
				continue
			}
		}
	}
	return replacedSlice
}

func ReplaceArgWithValueInString(arg string, args []CommandArg) string {
	for _, a := range args {
		argStart := strings.Index(arg, "{")
		argEnd := strings.Index(arg, "}")

		var extractedArg string

		if argStart != -1 && argEnd != -1 {
			extractedArg = arg[argStart:argEnd] + "}"
		} else {
			extractedArg = arg
		}

		if extractedArg == "{"+a.Name+"}" {
			return arg[:argStart] + a.Value + arg[argEnd+1:]
		}
	}

	return arg
}

func PrintError(msg string, urgent bool) {
	if urgent {
		fmt.Println(color.InBlackOverRed(msg))
		AppendToLogFile("error", " [ERROR] - "+msg)
		return
	}

	fmt.Println(color.InRed(msg))
	AppendToLogFile("error", msg)
}
func PrintInfo(msg string) {
	fmt.Println(color.InBlue(msg))
	AppendToLogFile("general", " [INFO] - "+msg)
}
func PrintAction(msg string) {
	fmt.Println(color.InGreen(msg))
	AppendToLogFile("general", " [ACTION] - "+msg)
}

func PrintWarning(msg string) {
	fmt.Println(color.InBlackOverYellow(msg))
	AppendToLogFile("general", " [WARNING] - "+msg)
}
