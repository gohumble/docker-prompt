package docker

import (
	"github.com/c-bata/go-prompt"
)

var commands = []prompt.Suggest{
	{Text: "ps", Description: "List containers"},
	{Text: "build", Description: "Build an image from a Dockerfile"},
	{Text: "images", Description: "List Images"},
	{Text: "exec", Description: "Run a command in a running container"},
	{Text: "run", Description: "Run a command in a new container"},
}

func (c *Completer) argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "build":

	case "exec":
		if len(args) == 2 {
			return getExecFirstOptionSuggestions()
		}
		if len(args) == 3 {
			return getPsSuggestions()
		}
		if len(args) == 4 {
			return getExecSampleCommandsSuggestions()
		}
	case "run":
		if len(args) == 2 {
			return c.getImagesSuggestions()
		}

	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
