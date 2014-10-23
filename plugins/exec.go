package plugins

import (
	"os"
	"os/exec"
	"path/filepath"
)

func (p *Plugins) node(args ...string) {
	bin := filepath.Join(p.nodePath, "bin", "node")
	runCmd(bin, args...)
}

func (p *Plugins) npm(args ...string) {
	bin := filepath.Join(p.nodePath, "bin", "npm")
	runCmd(bin, args...)
}

func runCmd(bin string, args ...string) {
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
