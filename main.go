package main

import (
	"os"

	"github.com/dickeyxxx/gonpm/context"
	"github.com/dickeyxxx/gonpm/plugins"
)

var ctx *context.Context
var topics []Topic

func init() {
	ctx = context.Parse(os.Args[1:]...)
	topics = []Topic{
		&plugins.Plugins{Context: ctx},
	}
}

func main() {
	topic := FindTopicByName(ctx.Topic)
	if topic == nil {
		Help()
		ctx.Exit(2)
	}
	topic.Run()
}
