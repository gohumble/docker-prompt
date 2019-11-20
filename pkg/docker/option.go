package docker

import "github.com/c-bata/go-prompt"

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
	case "exec":
		suggests = execOptions
	}
	return suggests
}
