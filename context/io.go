package context

import (
	"fmt"
	"os"
	"strings"
)

func (c *Context) Exit(code int) {
	c.exitFn(code)
}

func (c *Context) Stderrf(format string, a ...interface{}) {
	fmt.Fprintf(c.stderr, format, a...)
}

func (c *Context) Stderrln(a ...interface{}) {
	fmt.Fprintln(c.stderr, a...)
}

func (c *Context) Stdoutf(format string, a ...interface{}) {
	fmt.Fprintf(c.stdout, format, a...)
}

func (c *Context) Stdoutln(a ...interface{}) {
	fmt.Fprintln(c.stdout, a...)
}

func (c *Context) Logln(a ...interface{}) {
	c.logger.Println(a...)
	if debugging {
		fmt.Fprintln(c.stderr, a...)
	}
}

func (c *Context) Logf(format string, a ...interface{}) {
	c.logger.Printf(format, a...)
	if debugging {
		fmt.Fprintf(c.stderr, format, a...)
	}
}

var debugging = isDebugging()

func isDebugging() bool {
	debug := strings.ToUpper(os.Getenv("DEBUG"))
	if debug == "TRUE" || debug == "1" {
		return true
	}
	return false
}
