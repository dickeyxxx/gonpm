package plugins

import (
	"os"
	"os/user"
	"runtime"
)

func homeDir() string {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	return user.HomeDir + "/.gonpm"
}

func fileExists(path string) (bool, error) {
	var err error
	if runtime.GOOS == "windows" {
		// Windows doesn't seem to like using os.Stat
		_, err = os.Open(path)
	} else {
		_, err = os.Stat(path)
	}
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
