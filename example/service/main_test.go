package main_test

import (
	"github.com/LUSHDigital/core-mage/env"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	env.MustLoadDefaultTest(m)
	os.Exit(m.Run())
}

func Test(t *testing.T) {
	u := os.Getenv("MIGRATIONS_URL")
	switch u {
	case
		"file:///service/service/database/migrations",
		"file://service/database/migrations":
	default:
		t.Errorf("MIGRATIONS_URL was not suppoed to be %q", u)
	}
}
