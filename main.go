package main

import (
	"os"
	"runtime/debug"
	"time"

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
	run(topic)
}

func run(topic Topic) {
	ctx.Logf("Running %s:%s %s\n", ctx.Topic, ctx.Command, ctx.Args)
	before := time.Now()
	topic.Run()
	ctx.Logf("Finished in %s\n", (time.Since(before)))
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
