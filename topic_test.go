package main

import (
	"github.com/dickeyxxx/gonpm/context"
)

type fakeTopic struct {
	*context.Context
}

func (*fakeTopic) Name() string {
	return "faketopic"
}

func (f *fakeTopic) Help() {
	f.Stderrln("this is the help for faketopic")
}

func (f *fakeTopic) Run() {
	f.Stdoutln("faketopic has run")
}
