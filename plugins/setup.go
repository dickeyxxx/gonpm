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
	p.Stderrf("Setting up plugins... ")
	p.Logln("Creating plugins directory")
	err := os.MkdirAll(p.AppDir, 0777)
	must(err)
	p.Logln("Downloading node from", NODE_URL)
	resp, err := http.Get(NODE_URL)
	must(err)
	defer resp.Body.Close()
	uncompressed, err := gzip.NewReader(resp.Body)
	must(err)
	p.Logln("Extracting node to", p.nodePath)
	archive := tar.NewReader(uncompressed)
	for {
		hdr, err := archive.Next()
		if err == io.EOF {
			break
		}
		must(err)
		path := filepath.Join(p.AppDir, hdr.Name)
		switch {
		case hdr.FileInfo().IsDir():
			err = os.Mkdir(path, 0777)
			must(err)
		case hdr.Linkname != "":
			err = os.Symlink(hdr.Linkname, path)
			must(err)
		default:
			file, err := os.Create(path)
			must(err)
			defer file.Close()
			_, err = io.Copy(file, archive)
			must(err)
		}
		err = os.Chmod(path, hdr.FileInfo().Mode())
		must(err)
	}
	p.Logln("Finished installing node")
	p.Stderrln("done")
}
