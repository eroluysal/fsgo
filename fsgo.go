package main

import (
	"os"
	"path/filepath"
	"strconv"
)

func fNameAndExt(filename string) (name, ext string) {
	e := filepath.Ext(filename)
	n := filename[0 : len(filename)-len(e)]
	return n, e
}

func uniqueFInPath(name, path string) (fName string) {
	f, ext := fNameAndExt(name)
	for i := 0; ; i++ {
		if i > 0 {
			// If attempt is greater than 0 then append to attempt count to
			// file end.
			name = f + "-" + strconv.Itoa(i) + ext
		}
		// Check file validity. If file already does not exists then return the
		// name else continue loop until find available unique name.
		if _, err := os.Stat(filepath.Join(path, name)); os.IsNotExist(err) {
			return name
		}
	}
}
