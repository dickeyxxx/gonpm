package main

import (
	"fmt"
	"os"

	"github.com/dickeyxxx/gonpm/context"
)

func Help(ctx *context.Context) {
	var topic Topic
	if ctx != nil && len(ctx.Args) > 0 {
		topic = FindTopicByName(ctx.Args[0])
	}
	fmt.Printf("USAGE: %s\n", os.Args[0])
	if topic != nil {
		fmt.Printf("TOPIC: %s\n", topic.Name())
	}
}
