package targets

import (
	"io/ioutil"
	"os"
)

const gitlabCIFile = `---
include:
  - project: 'LUSHDigital/devops/gitlab-ci-includes'
    ref: master
    file: 'soa/common.yaml'

variables:
  TYPE: golang
`

func writeGitlabCIFile() error {
	raw := []byte(gitlabCIFile)
	return ioutil.WriteFile(".gitlab-ci.yml", raw, os.ModeDir|0664)
}
