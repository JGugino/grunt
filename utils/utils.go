package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/TwiN/go-color"
)

type CommandArg struct {
	Name  string
	Value string
}

type CommandReturn struct {
	Output []byte
	Err    error
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
			start, stop, extractedArg := ExtractArgFromSyntax(s)

			if extractedArg == "{"+a.Name+"}" {
				replacedSlice = append(replacedSlice, s[:start]+a.Value+s[stop+1:])
				continue
			}
		}
	}
	return replacedSlice
}

func ReplaceArgWithValueInString(arg string, args []CommandArg) string {
	for _, a := range args {
		start, stop, extractedArg := ExtractArgFromSyntax(arg)

		if extractedArg == "{"+a.Name+"}" {
			return arg[:start] + a.Value + arg[stop+1:]
		}
	}

	return arg
}

func ExtractArgFromSyntax(arg string) (start int, stop int, extracted string) {
	var extractedArg string

	argStart := strings.Index(arg, "{")
	argEnd := strings.Index(arg, "}")

	if argStart != -1 && argEnd != -1 {
		extractedArg = arg[argStart:argEnd] + "}"
	} else {
		extractedArg = arg
	}

	return argStart, argEnd, extractedArg
}

func ExtractPathOrReturnOriginalContent(content string) (hasPath bool, found string) {
	if i := strings.Index(content, "--path"); i == -1 {
		return false, content
	}

	_, _, contentPath := ExtractArgFromSyntax(content)

	fmt.Println(contentPath[1 : len(contentPath)-1])

	return true, contentPath[1 : len(contentPath)-1]
}

func ExecuteCommand(channel chan CommandReturn, command string, args []string) {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()

	commandReturn := CommandReturn{
		Output: output,
		Err:    err,
	}

	channel <- commandReturn
}

// Handles errors and has the ability to exit the program if fatal
func HandleError(err error, fatal bool) {
	if err != nil {
		if err.Error() == "no-create-name" {
			PrintWarningAndLog("You must provide a name for the config")
			os.Exit(1)
		} else if err.Error() == "invalid-log-type" {
			PrintErrorAndLog("Invalid log type (general or error)", false)
			os.Exit(1)
		} else if err.Error() == "invalid-log-args" {
			PrintErrorAndLog("Invalid log args", false)
			os.Exit(1)
		}

		PrintErrorAndLog(err.Error(), false)
		if fatal {
			os.Exit(1)
		}
	}
}

// Prints an error in black over a red background
func PrintError(msg string, urgent bool) {
	if urgent {
		fmt.Println(color.InBlackOverRed(msg))
		return
	}

	fmt.Println(color.InRed(msg))
}

// Prints program info in blue
func PrintInfo(msg string) {
	fmt.Println(color.InBlue(msg))
}

// Prints program action in green
func PrintAction(msg string) {
	fmt.Println(color.InGreen(msg))
}

// Prints program warning in black on a yellow background
func PrintWarning(msg string) {
	fmt.Println(color.InBlackOverYellow(msg))
}
