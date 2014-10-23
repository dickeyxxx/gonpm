package main

import (
	"testing"

	"github.com/dickeyxxx/gonpm/context"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHelp(t *testing.T) {
	Convey("it shows the USAGE message", t, func() {
		ctx = context.Parse()
		c := ctx.Mock()
		help()
		So(c.Stderr(), ShouldStartWith, "USAGE:")
	})

	Convey("with a topic", t, func() {
		ctx = context.Parse("help", "faketopic")
		topics = []Topic{&fakeTopic{ctx}}
		Convey("it shows help for that topic", func() {
			c := ctx.Mock()
			help()
			So(c.Stderr(), ShouldContainSubstring, "this is the help for faketopic")
		})
	})
}
