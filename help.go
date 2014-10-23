package main

import (
	"os"

	"github.com/dickeyxxx/gonpm/cli"
)

func help(command string, args ...string) {
	var topic *cli.Topic
	cli.Stderrf("USAGE: %s\n", os.Args[0])
	if len(args) > 0 {
		topic = topicByName(args[0])
		if topic != nil {
			topic.Help(command, args...)
			return
		}
	}
	for _, topic := range topics {
		cli.Stderrln(topic.Name)
	}
}
