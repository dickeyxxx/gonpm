package context

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Context struct {
	Topic, Command string
	Args           []string
	exitFn         func(code int)
	stdout         io.Writer
	stderr         io.Writer
}

func Parse(args ...string) *Context {
	ctx := &Context{
		exitFn: os.Exit,
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	if len(args) == 0 {
		return ctx
	}
	tc := strings.SplitN(args[0], ":", 2)
	ctx.Topic = tc[0]
	if len(tc) == 2 {
		ctx.Command = tc[1]
	}
	ctx.Args = args[1:]
	return ctx
}

func (c *Context) Exit(code int) {
	c.exitFn(code)
}

func (c *Context) Stderrf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.stderr, format, a...)
}

func (c *Context) Stderrln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.stderr, a...)
}

func (c *Context) Stdoutf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.stdout, format, a...)
}

func (c *Context) Stdoutln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.stdout, a...)
}
