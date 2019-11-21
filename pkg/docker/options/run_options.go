package options

import "github.com/c-bata/go-prompt"

var Run = []prompt.Suggest{
	{Text: "-d", Description: "Detached"},
	{Text: "-it", Description: "nteractive"},
	{Text: "--name", Description: "Identify container by name"},
	{Text: "-rm", Description: "Clean up"},
}
