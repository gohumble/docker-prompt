package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/gohumble/docker-promt/pkg/docker"
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
		prompt.OptionTitle("docker-prompt: interactive kubernetes client"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.Turquoise),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
	)
	p.Run()
}
