package plugins

import "path/filepath"

func (p *Plugins) node() string {
	return filepath.Join(p.nodePath, "bin", "node")
}

func (p *Plugins) npm() string {
	return filepath.Join(p.nodePath, "bin", "npm")
}
