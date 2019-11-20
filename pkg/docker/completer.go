package docker

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"os/exec"
	"strings"
)

func NewCompleter() (*Completer, error) {
	cmd := exec.Command("/bin/sh", "-c", "docker images")
	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil, err
	}
	imageRawlines := strings.Split(string(out), "\n")
	images := make([]Image, 0)
	for index, value := range imageRawlines {
		if index == 0 {
			continue
		}
		i := Image{}
		imagesRawCol := strings.Fields(value)
		if len(imagesRawCol) > 0 {
			i.repository = imagesRawCol[0]
			i.tag = imagesRawCol[1]
			i.imageId = imagesRawCol[2]
			i.size = imagesRawCol[6]
		}
		images = append(images, i)

	}
	return &Completer{
		images: images,
	}, nil
}

func (c *Completer) getImagesSuggestions() []prompt.Suggest {
	s := make([]prompt.Suggest, len(c.images))
	for i := range c.images {
		s[i] = prompt.Suggest{
			Text: c.images[i].imageId, Description: c.images[0].repository,
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
func getPsSuggestions() []prompt.Suggest {
	cmd := exec.Command("/bin/sh", "-c", "docker ps ")
	out, err := cmd.CombinedOutput()

	if err != nil {
		return nil
	}
	psRawlines := strings.Split(string(out), "\n")
	ps := make([]Ps, 0)
	for index, value := range psRawlines {
		if index == 0 {
			continue
		}
		i := Ps{}
		psRawCol := strings.Fields(value)
		if len(psRawCol) > 0 {
			i.containerId = psRawCol[0]
			i.image = psRawCol[1]
			i.names = psRawCol[10] + fmt.Sprintf(" | created %s %s ago", psRawCol[3], psRawCol[4])
		}
		ps = append(ps, i)

	}
	s := make([]prompt.Suggest, len(ps))
	for i := range ps {
		s[i] = prompt.Suggest{
			Text: ps[i].containerId, Description: ps[i].names,
		}
	}
	return s
}

type Completer struct {
	images []Image
	ps     []Ps
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

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	w := d.GetWordBeforeCursor()
	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, "-") {
		return optionCompleter(args, strings.HasPrefix(w, "--"))
	}
	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	return c.argumentsCompleter(args)
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
