package targets

import (
	"os"
	"path"
)

var (
	// ProjectType configures type of the project.
	ProjectType = "service"

	// ProjectName configures the name of the project.
	ProjectName = "service"

	// Environment describes the environment variables that should be sent with the target.
	Environment = CMDEnv{
		"GOPATH":    os.Getenv("GOPATH"),
		"GOPROXY":   os.Getenv("GOPROXY"),
		"GOMODPATH": os.Getenv("GOMODPATH"),
		"PWD":       os.Getenv("PWD"),
	}
)

// CMDEnv is used to wrap the command environment with convenience methods.
type CMDEnv map[string]string

// GoModPath derives the go module path from the environment.
func (e CMDEnv) GoModPath() string {
	var mod string
	if p := Environment["GOMODPATH"]; p != "" {
		mod = p
	}
	if p := Environment["GOPATH"]; p != "" && mod == "" {
		mod = path.Join(p, "mod")
	}
	if mod == "" {
		mod = "/go/pkg/mod"
	}
	return mod
}
