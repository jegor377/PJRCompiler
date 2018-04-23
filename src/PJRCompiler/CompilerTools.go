package PJRCompiler

import "strings"

func splitCommandAndParameters(command string) []string {
	var splitedCommand []string
	if strings.Count(command, "\"") > 0 {
		tmpSplitedCommand := strings.Split(command, "\"")
		for _, v := range(tmpSplitedCommand) {
			val := strings.TrimSpace(v)
			if len(val) != 0 {
				splitedCommand = append(splitedCommand, val)
			}
		}
		return splitedCommand
	}
	splitedCommand = strings.Split(command, " ")
	return splitedCommand
}