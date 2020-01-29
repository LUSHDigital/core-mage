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
