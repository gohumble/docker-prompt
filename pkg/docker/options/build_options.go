package options

import "github.com/c-bata/go-prompt"

var Build = []prompt.Suggest{
	{Text: "--file", Description: "Name of the Dockerfile (Default is ‘PATH/Dockerfile’)"},
	{Text: "--add-host", Description: "Add a custom host-to-IP mapping (host:ip)"},
	{Text: "--rm", Description: "Remove intermediate containers after a successful build"},
}
