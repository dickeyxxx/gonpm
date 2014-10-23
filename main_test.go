package main

import (
	"testing"

	"github.com/dickeyxxx/gonpm/context"
	. "github.com/smartystreets/goconvey/convey"
)

type fakeTopic struct {
	*context.Context
}

func (*fakeTopic) Name() string {
	return "faketopic"
}

func (*fakeTopic) Help() {
}

func (f *fakeTopic) Run() {
	f.Stdoutln("foobar")
}

func TestMain(t *testing.T) {
	Convey("with no topic it exits with code 2", t, func() {
		ctx = context.Parse()
		ctx.Mock()
		So(func() {
			main()
		}, ShouldPanicWith, 2)
	})

	Convey("with an invalid topic it exits with code 2", t, func() {
		ctx = context.Parse("myinvalidtopic")
		ctx.Mock()
		So(func() {
			main()
		}, ShouldPanicWith, 2)
	})

	Convey("with a valid topic it runs it", t, func() {
		ctx = context.Parse("faketopic")
		topics = []Topic{&fakeTopic{ctx}}
		mock := ctx.Mock()
		main()
		So(mock.Stdout(), ShouldEqual, "foobar\n")
	})
}