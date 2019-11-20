package docker

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var cmd *exec.Cmd

func Executor(s string) {
	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}
	//	if strings.Contains(s, "exec") {
	//		fields := strings.Fields(s)
	//		cmd = exec.Command("/bin/sh", "-c", "docker "+fields[0]+" -it "+fields[1]+" /bin/sh")
	//	} else {
	cmd = exec.Command("/bin/sh", "-c", "docker "+s)
	//	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}
