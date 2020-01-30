package mod_test

import (
	"os"
	"strings"
	"testing"

	"github.com/LUSHDigital/core-mage/env/internal/mod"
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
	if trim != "/env/internal/mod" {
		t.Errorf("\nworkdir: %s\n actual: %s", wd, dir)
		t.Errorf("\ntrimmed: %s", trim)
	}
}
