// Code generated by 'option-gen'. DO NOT EDIT.

package options

import (
	prompt "github.com/c-bata/go-prompt"
)

var Exec = []prompt.Suggest{
	prompt.Suggest{Text: "-d", Description: "Detached mode: run command in the background"},
	prompt.Suggest{Text: "--detach", Description: "Detached mode: run command in the background"},
	prompt.Suggest{Text: "-e", Description: "Set environment variables"},
	prompt.Suggest{Text: "--env list", Description: "Set environment variables"},
	prompt.Suggest{Text: "-i", Description: "Keep STDIN open even if not attached"},
	prompt.Suggest{Text: "--interactive", Description: "Keep STDIN open even if not attached"},
	prompt.Suggest{Text: "--privileged", Description: "Give extended privileges to the command"},
	prompt.Suggest{Text: "-t", Description: "Allocate a pseudo-TTY"},
	prompt.Suggest{Text: "--tty", Description: "Allocate a pseudo-TTY"},
	prompt.Suggest{Text: "-u", Description: "Username or UID (format: <name|uid>[:<group|gid>])"},
	prompt.Suggest{Text: "--user string", Description: "Username or UID (format: <name|uid>[:<group|gid>])"},
	prompt.Suggest{Text: "-w", Description: "Working directory inside the container"},
	prompt.Suggest{Text: "--workdir string", Description: "Working directory inside the container"},
}
