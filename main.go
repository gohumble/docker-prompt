package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	c "github.com/c-bata/go-prompt/completer"
	"github.com/gohumble/docker-prompt/pkg/docker"
)

var (
	version = "0"
)

func main() {
	completer, err := docker.NewCompleter()
	if err != nil {
		panic(err)
	}
	fmt.Printf("docker-prompt %s\n", version)
	fmt.Println("Please use `exit` to exit this program.")
	defer fmt.Println("Bye!")
	p := prompt.New(
		docker.Executor,
		completer.Do,
		prompt.OptionTitle("docker-prompt: interactive docker shell"),
		prompt.OptionPrefix("§ "),
		prompt.OptionInputTextColor(prompt.DefaultColor),
		prompt.OptionCompletionWordSeparator(c.FilePathCompletionSeparator),
	)
	p.Run()
}
