package docker

import (
	"github.com/c-bata/go-prompt"
	"github.com/docker/docker/api/types/registry"
)

var commands = []prompt.Suggest{
	{Text: "ps", Description: "List containers"},
	{Text: "build", Description: "Build an image from a Dockerfile"},
	{Text: "images", Description: "List Images"},
	{Text: "exec", Description: "Run a command in a running container"},
	{Text: "run", Description: "Run a command in a new container"},
	{Text: "info", Description: "Display system-wide information"},
	{Text: "cp", Description: "Copy files/folders between a container and the local filesystem"},
}

func seachResultArguments(result []registry.SearchResult) []prompt.Suggest {
	s := make([]prompt.Suggest, 0)
	for _, value := range result {
		s = append(s, prompt.Suggest{
			Text: value.Name, Description: value.Description,
		})
	}
	return s
}
func (c *Completer) argumentsCompleter(args []string, d prompt.Document) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "build":
		if len(args) == 2 {
			return optionCompleter(args, false)
		}
	case "exec":
		if len(args) == 2 {
			return optionCompleter(args, false)
		}
		if len(args) == 3 {
			return c.getContainerSuggestions()
		}
		if len(args) == 4 {
			return getExecSampleCommandsSuggestions()
		}
	case "run":
		// nice little search function here

		if len(args) >= 2 {

			sug := c.getImagesSuggestions(args[1])
			if len(sug) > 0 {
				c.Suggestions = sug
			}

			if len(args[len(args)-1]) > 1 && len(c.dockerWatcher.searchResultChan) > 0 {
				searchResult := <-c.dockerWatcher.searchResultChan
				c.Suggestions = seachResultArguments(searchResult)

			}
			return c.Suggestions
		}

	case "info":
		if len(args) == 2 {
			return optionCompleter(args, false)
		}
	case "cp":
		if len(args) == 2 {
			return optionCompleter(args, false)
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
