package context

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {
	Convey("with no arguments", t, func() {
		Convey("it returns an empty context", func() {
			ctx := Parse()
			So(ctx.Topic, ShouldBeBlank)
			So(ctx.Command, ShouldBeBlank)
			So(ctx.Args, ShouldBeEmpty)
		})
	})

	Convey("with `apps`", t, func() {
		Convey("it has a topic only", func() {
			ctx := Parse("apps")
			So(ctx.Topic, ShouldEqual, "apps")
			So(ctx.Command, ShouldBeBlank)
			So(ctx.Args, ShouldBeEmpty)
		})
	})

	Convey("with `apps:create`", t, func() {
		Convey("it has a topic and command", func() {
			ctx := Parse("apps:create")
			So(ctx.Topic, ShouldEqual, "apps")
			So(ctx.Command, ShouldEqual, "create")
			So(ctx.Args, ShouldBeEmpty)
		})
	})

	Convey("with `apps:create mynewapp`", t, func() {
		Convey("it has a topic, command and 1 argument", func() {
			ctx := Parse("apps:create", "mynewapp")
			So(ctx.Topic, ShouldEqual, "apps")
			So(ctx.Command, ShouldEqual, "create")
			So(ctx.Args[0], ShouldEqual, "mynewapp")
		})
	})

	Convey("with `apps:create mynewapp foobar`", t, func() {
		Convey("it has a topic, command and 2 arguments", func() {
			ctx := Parse("apps:create", "mynewapp", "foobar")
			So(ctx.Topic, ShouldEqual, "apps")
			So(ctx.Command, ShouldEqual, "create")
			So(ctx.Args[0], ShouldEqual, "mynewapp")
			So(ctx.Args[1], ShouldEqual, "foobar")
		})
	})
}
