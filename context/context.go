package context

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Context struct {
	Topic, Command string
	Args           []string
	AppDir         string
	logger         *log.Logger
	Stdout         io.Writer
	Stderr         io.Writer
	exitFn         func(code int)
}

func Parse(args ...string) *Context {
	ctx := &Context{
		AppDir: homeDir() + "/.gonpm",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		exitFn: os.Exit,
	}
	ctx.logger = newLogger(ctx.AppDir + "/gonpm.log")
	ctx.Topic, ctx.Command, ctx.Args = parse(args...)
	return ctx
}

func newLogger(path string) *log.Logger {
	err := os.MkdirAll(filepath.Dir(path), 0777)
	must(err)
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	must(err)
	return log.New(file, "", log.LstdFlags)
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

func must(err error) {
	if err != nil {
		panic(err)
	}
}
