package main

type Topic interface {
	Name() string
	Run()
	Help()
}

func FindTopicByName(name string) Topic {
	for _, topic := range topics {
		if name == topic.Name() {
			return topic
		}
	}
	return nil
}
