package main

import (
	"os"

	"github.com/dickeyxxx/gonpm/cli"
)

func help(command string, args ...string) {
	var topic *cli.Topic
	cli.Stderrf("USAGE: %s COMMAND [--app APP] [command-specific-options]\n\n", os.Args[0])

	cli.Stderrf("Help topics, type \"%s help TOPIC\" for more details:\n\n", os.Args[0])
	if len(args) > 0 {
		topic = topicByName(args[0])
		if topic != nil {
			topic.Help(command, args...)
			return
		}
	}
	for _, topic := range topics {
		cli.Stderrf("  %s \t# %s\n", topic.Name, topic.Description)
	}
}
