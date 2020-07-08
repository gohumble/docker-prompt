package docker

import (
	"github.com/c-bata/go-prompt"
	"github.com/docker/docker/api/types/registry"
	"github.com/gohumble/docker-prompt/pkg/docker/options"
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

func searchResultArguments(result []registry.SearchResult) []prompt.Suggest {
	s := make([]prompt.Suggest, 0)
	for _, value := range result {
		s = append(s, prompt.Suggest{
			Text: value.Name, Description: value.Description,
		})
	}
	return s
}

func (c *Completer) CommandCompleter(args []string, d prompt.Document, currentWord string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}
	first := args[0]
	switch first {
	case "build":
		if len(args) == 2 {
			return options.Completer(args, false)
		}
	case "exec":
		if len(args) == 2 {
			return options.Completer(args, false)
		}
		if len(args) == 3 {
			return c.getContainerSuggestions()
		}
		if len(args) == 4 {
			return getExecSampleCommandsSuggestions()
		}
	case "run":
		// nice little search function here
		if len(args) > 3 {
			if len(currentWord) > 0 {
				if currentWord[:1] == "/" {
					return getTerminalSuggestions()
				}
			}
		}
		if len(args) >= 2 {
			sug := c.getImagesSuggestions(currentWord)
			if len(sug) > 0 {
				return sug
			} else {
				select {
				case searchResult := <-c.dockerWatcher.searchResultChan:
					c.Suggestions = searchResultArguments(searchResult)
				default:
					c.Suggestions = make([]prompt.Suggest, 0)
					c.Suggestions = append(c.Suggestions, prompt.Suggest{Text: "searching..."})
				}
			}
			return c.Suggestions
		}
	case "info":
		if len(args) == 2 {
			return options.Completer(args, false)
		}
	case "cp":
		if len(args) == 2 {
			return options.Completer(args, false)
		}
	default:
		return []prompt.Suggest{}
	}
	return []prompt.Suggest{}
}
