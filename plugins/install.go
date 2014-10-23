package plugins

import (
	"io"
	"os/exec"
)

func (p *Plugins) install() {
	if len(p.Args) != 1 {
		p.Help()
		p.Exit(2)
	}
	name := p.Args[0]
	p.Stderrf("Installing plugin %s...\n", name)
	cmd := exec.Command(p.npm(), "install", "-g", "hk-plugin")
	stderr, err := cmd.StderrPipe()
	must(err)
	err = cmd.Start()
	go io.Copy(p.Stderr, stderr)
	must(err)
	err = cmd.Wait()
	must(err)
}
