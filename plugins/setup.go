package plugins

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func (p *Plugins) Setup() {
	if exists, _ := fileExists(p.nodePath); exists == true {
		return
	}
	p.Stderrln("Setting up plugins... ")
	path := filepath.Join(p.AppDir, "plugins")
	err := os.MkdirAll(path, 0777)
	must(err)
	resp, err := http.Get(NODE_URL)
	must(err)
	defer resp.Body.Close()
	uncompressed, err := gzip.NewReader(resp.Body)
	must(err)
	archive := tar.NewReader(uncompressed)
	for {
		hdr, err := archive.Next()
		if err == io.EOF {
			break
		}
		must(err)
		path := filepath.Join(path, hdr.Name)
		if hdr.FileInfo().IsDir() {
			err = os.Mkdir(path, 0777)
			must(err)
		} else {
			file, err := os.Create(path)
			must(err)
			defer file.Close()
			_, err = io.Copy(file, archive)
			must(err)
		}
	}
	err = os.Chmod(filepath.Join(p.nodePath, "bin", "node"), 0777)
	must(err)
	err = os.Chmod(filepath.Join(p.nodePath, "bin", "npm"), 0777)
	must(err)
}
