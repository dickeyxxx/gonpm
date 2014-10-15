package main

import (
	"fmt"
	"log"

	"github.com/dickeyxxx/gonpm/plugins"
)

func main() {
	if err := plugins.Setup(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("starting")
	plugins.ExecNode("-v")
	fmt.Println("done")
}
