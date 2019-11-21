package docker

import (
	"github.com/c-bata/go-prompt"
	"github.com/gohumble/docker-promt/pkg/docker/options"
)

var optionHelp = []prompt.Suggest{
	{Text: "-h"},
	{Text: "--help"},
}

func optionCompleter(args []string, long bool) []prompt.Suggest {
	l := len(args)
	if l <= 1 {
		if long {
			return prompt.FilterHasPrefix(optionHelp, "--", false)
		}
		return optionHelp
	}

	var suggests []prompt.Suggest
	commandArgs, _ := excludeOptions(args)
	switch commandArgs[0] {
	case "build":
		suggests = options.Build
	case "exec":
		suggests = options.Exec
	case "run":
		suggests = options.Run
	case "info":
		suggests = options.Info
	case "cp":
		suggests = options.Cp

	}
	return suggests
}
