package options

import "github.com/c-bata/go-prompt"

var Info = []prompt.Suggest{
	{Text: "--format", Description: "Format the output using the given Go template ('{{json .}}')"},
}
