package mod

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ModulePath will return the module path of the current directory and return an error if the project is not within a module.
func ModulePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return dir, err
	}
	return backtrack(dir, "go.mod")
}

func search(dir, fname string) (bool, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, f := range files {
		if f.IsDir() || f.Name() != fname {
			continue
		}
		return true, nil
	}
	return false, nil
}

func backtrack(dir, fname string) (string, error) {
	if !filepath.IsAbs(dir) {
		return dir, fmt.Errorf("dir must be an absolute file path")
	}
	for {
		found, err := search(dir, fname)
		if err != nil {
			return dir, err
		}
		if !found {
			if dir == "/" {
				return dir, fmt.Errorf("cannot find the file %q in any subsequent directories", fname)
			}
			dir = filepath.Dir(dir)
			continue
		}
		return dir, nil
	}
}
