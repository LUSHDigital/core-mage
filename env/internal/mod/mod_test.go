package mod_test

import (
	"os"
	"strings"
	"testing"

	"github.com/LUSHDigital/core-mage/mod"
)

func TestModulePath(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := mod.ModulePath()
	if err != nil {
		t.Fatal(err)
	}
	trim := strings.TrimPrefix(wd, dir)
	if trim != "/mod" {
		t.Errorf("\nworkdir: %s\n  actual: %s", wd, dir)
	}
}
