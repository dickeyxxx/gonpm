package context

import (
	"io"
	"os"
	"strings"
)

type Context struct {
	Topic, Command string
	Args           []string
	AppDir         string
	exitFn         func(code int)
	stdout         io.Writer
	stderr         io.Writer
}

func Parse(args ...string) *Context {
	ctx := &Context{
		AppDir: homeDir() + "/.gonpm",
		exitFn: os.Exit,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	ctx.Topic, ctx.Command, ctx.Args = parse(args...)
	return ctx
}

func parse(input ...string) (topic, command string, args []string) {
	if len(input) == 0 {
		return
	}
	tc := strings.SplitN(input[0], ":", 2)
	topic = tc[0]
	if len(tc) == 2 {
		command = tc[1]
	}
	args = input[1:]
	return topic, command, args
}
