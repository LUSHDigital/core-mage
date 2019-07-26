package targets

import (
	"io/ioutil"
)

const gitIgnoreFile = `infra/local.env
vendor/
build/
data/
`

func writeGitIgnoreFile() error {
	raw := []byte(gitIgnoreFile)
	return ioutil.WriteFile(".gitignore", raw, 0664)
}
