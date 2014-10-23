package main

import "os"

func Help() {
	var topic Topic
	if len(ctx.Args) > 0 {
		topic = FindTopicByName(ctx.Args[0])
	}
	ctx.Stderrf("USAGE: %s\n", os.Args[0])
	if topic != nil {
		ctx.Stderrf("TOPIC: %s\n", topic.Name())
	}
}
