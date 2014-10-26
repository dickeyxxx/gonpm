package plugins

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func setup() {
	home, err := homedir.Dir()
	must(err)
	fmt.Printf("starting... [%s]\n", home)
	//cmd := exec.Command(filepath.Join(home, ".hk", "npm"), "ls")
	//cmd := exec.Command(filepath.Join(home, ".hk", "npm"), "install", "dickeyxxx/heroku-production-check")
	cmd := exec.Command(filepath.Join(home, ".hk", "node"), "-e", "console.log(require('heroku-production-check'))")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	must(err)
	fmt.Println("done")
}

func must(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
