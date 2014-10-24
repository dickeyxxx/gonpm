package plugins

import (
	"os/exec"

	"github.com/dickeyxxx/gonpm/cli"
)

type Plugin struct {
	*cli.Topic
	Version  string `json:"version"`
	From     string `json:"from"`
	Resolved string `json:"resolved"`
}

func pluginRun(name string) func(command string, args ...string) {
	return func(command string, args ...string) {
		runNode(`require('` + name + `').run()`)
	}
}

func pluginHelp(name string) func(command string, args ...string) {
	return func(command string, args ...string) {
		runNode(`require('` + name + `').help()`)
	}
}

func runNode(script string) {
	cmd := exec.Command(nodePath, "-e", script)
	cmd.Stdout = cli.Stdout
	cmd.Stderr = cli.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
