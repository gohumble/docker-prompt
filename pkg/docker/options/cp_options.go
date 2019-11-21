package options

import "github.com/c-bata/go-prompt"

var Cp = []prompt.Suggest{
	{Text: "-a", Description: "Archive mode (copy all uid/gid information)"},
	{Text: "-L", Description: "Always follow symbol link in SRC_PATH"},
}
