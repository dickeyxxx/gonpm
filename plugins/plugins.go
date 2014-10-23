package plugins

import "github.com/dickeyxxx/gonpm/context"

type Plugins struct {
	*context.Context
}

func (*Plugins) Name() string {
	return "plugins"
}

func (p *Plugins) Initialize(ctx *context.Context) {
	p.Context = ctx
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
