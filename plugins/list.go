package plugins

import (
	"encoding/json"
	"io"
	"os/exec"
)

type pkg struct {
	Name     string
	Version  string `json:"version"`
	From     string `json:"from"`
	Resolved string `json:"resolved"`
}

func (p *Plugins) list() {
	p.Logln("Listing plugins...")
	for _, pkg := range p.npmPackages() {
		p.Stdoutln(pkg.Name)
	}
}

func (p *Plugins) npmPackages() []pkg {
	cmd := exec.Command(p.npm(), "list", "--json")
	stderr, err := cmd.StderrPipe()
	must(err)
	stdout, err := cmd.StdoutPipe()
	must(err)
	err = cmd.Start()
	go io.Copy(p.Stderr, stderr)
	must(err)
	var doc map[string]map[string]pkg
	err = json.NewDecoder(stdout).Decode(&doc)
	must(err)
	err = cmd.Wait()
	must(err)
	var packages []pkg
	for name, p := range doc["dependencies"] {
		p.Name = name
		packages = append(packages, p)
	}
	return packages
}
