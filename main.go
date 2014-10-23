package main

import (
	"os"

	"github.com/dickeyxxx/gonpm/context"
	"github.com/dickeyxxx/gonpm/plugins"
)

var ctx *context.Context = context.Parse(os.Args[1:]...)
var topics []Topic = []Topic{
	&plugins.Plugins{},
}

func main() {
	initializeTopics()
	topic := topicByName(ctx.Topic)
	if topic == nil {
		help()
		ctx.Exit(2)
	}
	topic.Run()
}
