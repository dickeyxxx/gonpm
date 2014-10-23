package plugins

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const NODE_VERSION = "v0.10.32"

func nodeString() string {
	return "node-" + NODE_VERSION + "-darwin-x64"
}

func (p *Plugins) Setup() {
	if exists, _ := fileExists(p.nodePath()); exists == true {
		return
	}
	fmt.Println("Downloading " + nodeString() + "...")
	path := filepath.Join(p.AppDir, "plugins")
	err := os.MkdirAll(path, 0777)
	must(err)
	resp, err := http.Get("http://nodejs.org/dist/" + NODE_VERSION + "/" + nodeString() + ".tar.gz")
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
	err = os.Chmod(filepath.Join(p.nodePath(), "bin", "node"), 0777)
	if err != nil {
		panic(err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (p *Plugins) nodePath() string {
	return filepath.Join(p.AppDir, "plugins", nodeString())
}
