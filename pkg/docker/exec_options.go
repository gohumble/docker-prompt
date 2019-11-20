package docker

import "github.com/c-bata/go-prompt"

var execOptions = []prompt.Suggest{
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
