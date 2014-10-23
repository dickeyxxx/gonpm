package plugins

import (
	"os"
	"os/exec"

	"github.com/dickeyxxx/gonpm/context"
)

type Plugin struct {
	*Plugins
	pkg *pkg
}

func (p *Plugin) Name() string {
	return p.pkg.Name
}

func (p *Plugin) Initialize(ctx *context.Context) {
	p.Context = ctx
}

func (p *Plugin) Help() {
	p.Stderrln("Running", p.Name())
	script := `
	var plugin = require('` + p.Name() + `')
	plugin.help()
	`
	cmd := exec.Command(p.node(), "-e", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (p *Plugin) Run() {
	p.Stderrln("Running", p.Name())
	script := `
	var plugin = require('` + p.Name() + `')
	plugin()
	`
	cmd := exec.Command(p.node(), "-e", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (p *Plugins) LoadPluginTopics() []*Plugin {
	plugins := []*Plugin{}
	for _, pkg := range p.npmPackages() {
		plugin := &Plugin{p, &pkg}
		plugins = append(plugins, plugin)
	}
	return plugins
}
