package main

import (
	"os"
	"runtime/debug"

	"github.com/dickeyxxx/gonpm/context"
	"github.com/dickeyxxx/gonpm/plugins"
)

var ctx *context.Context = context.Parse(os.Args[1:]...)
var topics []Topic = []Topic{
	&plugins.Plugins{},
}

func main() {
	defer handlePanic()
	initializeTopics()
	topic := topicByName(ctx.Topic)
	if topic == nil {
		help()
		ctx.Exit(2)
	}
	ctx.Logf("Running %s:%s %s\n", ctx.Topic, ctx.Command, ctx.Args)
	topic.Run()
}

func handlePanic() {
	if e := recover(); e != nil {
		switch e := e.(type) {
		case int:
			// This is for when we stub out ctx.Exit
			panic(e)
		default:
			ctx.Logf("ERROR: %s\n%s", e, debug.Stack())
			ctx.Stderrln("ERROR:", e)
			ctx.Exit(1)
		}
	}
}
