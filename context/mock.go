package context

import "bytes"

type MockContext struct {
	*Context
	ExitCode int
}

func (c *Context) Mock() *MockContext {
	mock := &MockContext{Context: c}
	var stderr bytes.Buffer
	var stdout bytes.Buffer
	mock.stderr = &stderr
	mock.stdout = &stdout
	mock.exitFn = func(code int) {
		panic(code)
	}
	return mock
}

func (m *MockContext) Stderr() string {
	return m.stderr.(*bytes.Buffer).String()
}

func (m *MockContext) Stdout() string {
	return m.stdout.(*bytes.Buffer).String()
}
