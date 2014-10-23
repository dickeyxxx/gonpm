package plugins

func (p *Plugins) install() {
	if len(p.Args) != 1 {
		p.Help()
		p.Exit(2)
	}
	p.Stdoutln("installing...")
}