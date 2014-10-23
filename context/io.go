package context

import "fmt"

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
