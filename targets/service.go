package targets

import (
	"os"
)

var (
	// ServiceDir represents the directory used for the service go files.
	ServiceDir = "service"
)

func writeServiceDir() error {
	return os.MkdirAll(ServiceDir, os.ModePerm)
}
