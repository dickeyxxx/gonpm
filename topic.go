package main

import "github.com/dickeyxxx/gonpm/context"

type Topic interface {
	Name() string
	Initialize(ctx *context.Context)
	Run()
	Help()
}

func topicByName(name string) Topic {
	for _, topic := range topics {
		if name == topic.Name() {
			return topic
		}
	}
	return nil
}

func initializeTopics() {
	for _, topic := range topics {
		topic.Initialize(ctx)
	}
}
