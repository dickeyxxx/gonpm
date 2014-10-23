package context

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExit(t *testing.T) {
	Convey("with a stubbed exit handler", t, func() {
		Convey("it exits with code 88", func() {
			ctx := Parse().Mock()
			So(func() {
				ctx.Exit(88)
			}, ShouldPanicWith, 88)
		})
	})
}

func TestStderr(t *testing.T) {
	Convey("Stderrln shows someloginfo", t, func() {
		ctx := Parse().Mock()
		ctx.Stderrln("someloginfo")
		So(ctx.Stderr(), ShouldEqual, "someloginfo\n")
	})

	Convey("Stderrf shows someloginfo", t, func() {
		ctx := Parse().Mock()
		ctx.Stderrf("someloginfo")
		So(ctx.Stderr(), ShouldEqual, "someloginfo")
	})
}

func TestStdout(t *testing.T) {
	Convey("Stdoutln shows someloginfo", t, func() {
		ctx := Parse().Mock()
		ctx.Stdoutln("someloginfo")
		So(ctx.Stdout(), ShouldEqual, "someloginfo\n")
	})

	Convey("Stdoutf shows someloginfo", t, func() {
		ctx := Parse().Mock()
		ctx.Stdoutf("someloginfo")
		So(ctx.Stdout(), ShouldEqual, "someloginfo")
	})
}
