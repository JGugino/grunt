package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
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
