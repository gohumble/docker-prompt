// Code generated by 'option-gen'. DO NOT EDIT.

package options

import (
	prompt "github.com/c-bata/go-prompt"
)

var Events = []prompt.Suggest{
	prompt.Suggest{Text: "--since string", Description: "Show all events created since timestamp"},
	prompt.Suggest{Text: "--until string", Description: "Stream events until this timestamp"},
}
