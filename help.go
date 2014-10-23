package main

import "os"

func help() {
	var topic Topic
	if len(ctx.Args) > 0 {
		topic = topicByName(ctx.Args[0])
	}
	ctx.Stderrf("USAGE: %s\n", os.Args[0])
	for _, topic := range topics {
		ctx.Stderrln(topic.Name())
	}
	if topic != nil {
		topic.Help()
	}
}
