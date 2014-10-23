package plugins

import (
	"testing"

	"github.com/dickeyxxx/gonpm/context"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTopic(t *testing.T) {
	Convey("it shows the help", t, func() {
		ctx := context.Parse("plugins:install").Mock()
		plugins := &Plugins{Context: ctx.Context}
		So(func() {
			plugins.Run()
		}, ShouldPanicWith, 2)
		So(ctx.Stderr(), ShouldEqual, "plugins help for: install\n")
	})
}
