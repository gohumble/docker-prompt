package docker

import (
	"github.com/c-bata/go-prompt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

type Completer struct {
	Suggestions   []prompt.Suggest
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
		return []prompt.Suggest{}
	}

	args := strings.Split(d.TextBeforeCursor(), " ")

	w := d.GetWordBeforeCursor()
	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}
	if len(args) > 1 && len(w) > 1 {
		go c.dockerWatcher.Search(args[1])
	}
	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	return c.argumentsCompleter(args, d)
}

func (c Completer) getImagesSuggestions(filter string) []prompt.Suggest {
	s := make([]prompt.Suggest, 0)
	for _, value := range c.images {
		id := formatImageId(value.ID)
		desc := formatImageDescription(value)
		if strings.Contains(id, filter) {
			s = append(s, prompt.Suggest{
				Text: id, Description: desc,
			})
		}
	}
	return s
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

type Image struct {
	repository string
	tag        string
	imageId    string
	created    string
	size       string
}
type Ps struct {
	containerId string
	image       string
	names       string
}

func excludeOptions(args []string) ([]string, bool) {
	l := len(args)
	if l == 0 {
		return nil, false
	}
	cmd := args[0]
	filtered := make([]string, 0, l)

	var skipNextArg bool
	for i := 0; i < len(args); i++ {
		if skipNextArg {
			skipNextArg = false
			continue
		}

		if cmd == "logs" && args[i] == "-f" {
			continue
		}

		for _, s := range []string{
			"-f", "--filename",
			"-n", "--namespace",
			"-s", "--server",
			"--kubeconfig",
			"--cluster",
			"--user",
			"-o", "--output",
			"-c",
			"--container",
		} {
			if strings.HasPrefix(args[i], s) {
				if strings.Contains(args[i], "=") {
					// we can specify option value like '-o=json'
					skipNextArg = false
				} else {
					skipNextArg = true
				}
				continue
			}
		}
		if strings.HasPrefix(args[i], "-") {
			continue
		}

		filtered = append(filtered, args[i])
	}
	return filtered, skipNextArg
}
