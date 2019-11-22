package main

import (
	"fmt"
	"github.com/gohumble/docker-prompt/pkg/docker"
	"github.com/gohumble/go-prompt"
	"github.com/gohumble/go-prompt/completer"
)

var (
	version string = "0"
)

func main() {
	c, err := docker.NewCompleter()
	if err != nil {
		panic(err)
	}
	fmt.Printf("docker-prompt %s\n", version)
	fmt.Println("Please use `exit` to exit this program.")
	defer fmt.Println("Bye!")
	p := prompt.New(
		docker.Executor,
		c.Complete,
		prompt.OptionTitle("docker-prompt: interactive docker shell"),
		prompt.OptionPrefix("§ "),
		prompt.OptionInputTextColor(prompt.DefaultColor),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
