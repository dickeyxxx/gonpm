package plugins

import (
	"os"
	"path/filepath"

	"github.com/dickeyxxx/gonpm/context"
)

type Plugins struct {
	*context.Context
	nodePath string
}

func (*Plugins) Name() string {
	return "plugins"
}

func (p *Plugins) Initialize(ctx *context.Context) {
	p.Context = ctx
	p.nodePath = filepath.Join(p.AppDir, NODE_STRING)
	os.Setenv("NODE_PATH", filepath.Join(p.AppDir, "lib", "node_modules"))
	os.Setenv("NPM_CONFIG_GLOBAL", "true")
	os.Setenv("NPM_CONFIG_PREFIX", p.AppDir)
	p.Setup()
}

func (p *Plugins) Run() {
	switch p.Command {
	case "install":
		p.install()
	case "list":
	case "":
		p.list()
	default:
		p.Help()
	}
}

func (p *Plugins) Help() {
	p.Stderrln("plugins help for:", p.Command)
}
