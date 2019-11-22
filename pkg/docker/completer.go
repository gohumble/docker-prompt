package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gohumble/docker-prompt/pkg/docker/options"
	"github.com/gohumble/go-prompt"
	"strings"
)

type Completer struct {
	Suggestions   []prompt.Suggest
	LastWord      string
	docker        *client.Client
	dockerWatcher Watcher
	containers    []types.Container
	images        []types.ImageSummary
}

func NewCompleter() (*Completer, error) {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	completer := &Completer{
		dockerWatcher: NewWatcher(dockerClient),
		docker:        dockerClient,
		Suggestions:   make([]prompt.Suggest, 0),
	}

	completer.watchForChanges()
	return completer, nil
}

func (c *Completer) watchForChanges() {
	c.dockerWatcher.Start(c)
}
func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	cursorPrefix := d.TextBeforeCursor()
	if cursorPrefix == "" {
		c.Suggestions = []prompt.Suggest{}
		return c.Suggestions
	}

	args := strings.Split(d.TextBeforeCursor(), " ")

	w := d.GetWordBeforeCursor()
	if len(args) > 1 && len(w) > 1 {
		//		SearchGroup.Do(w, c.Search())
		go c.dockerWatcher.Search(w)
	}
	if c.LastWord == w {
		return c.Suggestions
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		c.Suggestions = options.Completer(args, strings.HasPrefix(w, "--"))
		return c.Suggestions
	}

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			c.Suggestions = []prompt.Suggest{}
			return c.Suggestions
		}
	}

	c.Suggestions = c.CommandCompleter(args, d, w)
	c.LastWord = w

	return c.Suggestions
}
func (c *Completer) Search() func() (result interface{}, err error) {
	return func() (interface{}, error) {
		go c.dockerWatcher.Search(c.LastWord)
		return nil, nil
	}
}

func (c Completer) getImagesSuggestions(filter string) []prompt.Suggest {
	s := make([]prompt.Suggest, 0)
	if !strings.Contains(filter, "-") {
		for _, value := range c.images {
			id := formatImageId(value.ID)
			desc := formatImageDescription(value)
			if strings.Contains(id, filter) {
				s = append(s, prompt.Suggest{
					Text: id, Description: desc,
				})
			}
		}
	}
	return s
}
func getTerminalSuggestions() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "/bin/sh", Description: "sh"},
		{Text: "/bin/bash", Description: "bash"},
	}
}
func getExecSampleCommandsSuggestions() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "/bin/sh", Description: "sh"},
		{Text: "echo HELLO WORLD", Description: "sample"},
	}
}
func getExecFirstOptionSuggestions() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "-it", Description: "lets go mode"},
		{Text: "-e", Description: "Set environment variables"},
		{Text: "-d", Description: "Detached mode: run command in the background"},
		{Text: "-privileged", Description: "Give extended privileges to the command"},
		{Text: "-u", Description: "Username or UID (format: <name|uid>[:<group|gid>])"},
		{Text: "-w", Description: "Working directory inside the container"},
		{Text: "--detach-keys", Description: "Override the key sequence for detaching a container"},
		{Text: "-t", Description: "Allocate a pseudo-TTY"},
		{Text: "-i", Description: "interactive mode"},
	}
}

func (c *Completer) getContainerSuggestions() []prompt.Suggest {
	s := make([]prompt.Suggest, len(c.containers))
	for key, value := range c.containers {
		s[key] = prompt.Suggest{
			Text: value.Image, Description: value.State,
		}
	}
	return s
}
