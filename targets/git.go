package targets

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/magefile/mage/sh"
)

const gitIgnoreFile = `infra/local.env
infra/personal.env
vendor/
build/
data/
modules/
`

var (
	// SubmodulesPath is used to describe what path to use for git submodules.
	SubmodulesPath = "modules/"
)

func initGit() error {
	return Exec(GitBin, "init")
}

func pullSubmodule(name, branch string) error {
	path := submodulePath(name)
	os.Chdir(path)
	defer os.Chdir(Environment["PWD"])
	if _, err := sh.Exec(Environment, os.Stdout, os.Stderr, GitBin, "checkout", branch); err != nil {
		return err
	}
	if _, err := sh.Exec(Environment, os.Stdout, os.Stderr, GitBin, "pull"); err != nil {
		return err
	}
	return nil
}

func submodulePath(name string) string {
	return path.Join(SubmodulesPath, name)
}

func addSubmodule(name, repo, branch string) error {
	path := submodulePath(name)
	return Exec(GitBin,
		"submodule", "-q",
		"add", "-f", "--name", name, "-b", branch, repo, path,
	)
}

func initSubmodule(name string) error {
	path := submodulePath(name)
	return Exec(GitBin, "submodule", "-q", "init", path)
}

func updateSubmodule(name string) error {
	path := submodulePath(name)
	return Exec(GitBin, "submodule", "-q", "update", "-f", "--checkout", path)
}

func deinitSubmodule(name string) error {
	path := submodulePath(name)
	return Exec(GitBin,
		"submodule", "-q",
		"deinit", "-f", path,
	)
}

func removeSubmodule(name string) error {
	path := submodulePath(name)
	return Exec(GitBin,
		"rm", "--cached", "--ignore-unmatch", "-q", path,
	)
}

func removeSubmoduleDir(name string) error {
	path := submodulePath(name)
	return Exec("rm", "-rf", path)
}

func writeGitIgnoreFile() error {
	raw := []byte(gitIgnoreFile)
	return ioutil.WriteFile(".gitignore", raw, 0664)
}
