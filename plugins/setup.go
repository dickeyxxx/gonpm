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

func Setup() error {
	if exists, _ := fileExists(nodePath()); exists == true {
		return nil
	}
	fmt.Println("Downloading " + nodeString() + "...")
	path := filepath.Join(homeDir(), "plugins")
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
	err = os.Chmod(filepath.Join(nodePath(), "bin", "node"), 0777)
	if err != nil {
		panic(err)
	}
	return nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func nodePath() string {
	return filepath.Join(homeDir(), "plugins", nodeString())
}
