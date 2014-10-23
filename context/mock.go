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
	mock.Stderr = &stderr
	mock.Stdout = &stdout
	mock.exitFn = func(code int) {
		panic(code)
	}
	return mock
}

func (m *MockContext) GetStderr() string {
	return m.Stderr.(*bytes.Buffer).String()
}

func (m *MockContext) GetStdout() string {
	return m.Stdout.(*bytes.Buffer).String()
}
