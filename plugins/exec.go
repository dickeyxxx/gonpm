package plugins

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func (p *Plugins) ExecNode(args ...string) {
	bin := filepath.Join(p.nodePath(), "bin", "node")
	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
