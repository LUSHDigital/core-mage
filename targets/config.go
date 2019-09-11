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

	// GitBin is the executable name of Git.
	GitBin = "git"

	// GoBin is the executable name of Go.
	GoBin = "go"

	// DockerBin is the executable name of Docker.
	DockerBin = "docker"

	// ComposeBin is the executable name of Docker Compose.
	ComposeBin = "docker-compose"

	// LocalHost is the IP or hostname of your local machine.
	LocalHost = "127.0.0.1"

	// MigrationsURLEnvVar is the name of the migrations URL environment variable.
	MigrationsURLEnvVar = "MIGRATIONS_URL"

	// MigrationsURLLocal is the environment variable for the migrations path in the local environment.
	MigrationsURLLocal = "file://service/database/migrations"

	// MigrationsURLDev is the environment variable for the migrations path in the development environment.
	MigrationsURLDev = "file:///migrations"

	// MigrationsURLTest is the environment variable for the migrations path in the test environment.
	MigrationsURLTest = "file:///service/service/database/migrations"
)

// Environment describes the environment variables that should be sent with the target.
var Environment = CMDEnv{
	"GOPATH":    os.Getenv("GOPATH"),
	"GOPROXY":   os.Getenv("GOPROXY"),
	"GOMODPATH": os.Getenv("GOMODPATH"),
	"PWD":       os.Getenv("PWD"),
}

// CMDEnv is used to wrap the command environment with convenience methods.
type CMDEnv map[string]string

// GoModPath derives the go module path from the environment.
func (e CMDEnv) GoModPath() string {
	var mod string
	if p := e["GOMODPATH"]; p != "" {
		mod = p
	}
	if p := e["GOPATH"]; p != "" && mod == "" {
		mod = path.Join(p, "mod")
	}
	if mod == "" {
		mod = "/go/pkg/mod"
	}
	return mod
}

// GoProxy will return the configured go proxy or provide a default.
func (e CMDEnv) GoProxy() string {
	const defaultGoProxy = "https://proxy.golang.org"
	if proxy := e["GOPROXY"]; proxy != "" {
		return proxy
	}
	return defaultGoProxy
}
