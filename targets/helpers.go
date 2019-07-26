package targets

import (
	"os"

	"github.com/magefile/mage/sh"
)

// Exec will execute a command with the default environment and print to standard output.
func Exec(bin string, args ...string) error {
	_, err := sh.Exec(Environment, os.Stdout, os.Stderr, bin, args...)
	return err
}
