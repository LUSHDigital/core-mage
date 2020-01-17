package env_test

import (
	"os"
	"testing"

	"github.com/LUSHDigital/core-mage/env"
	"github.com/LUSHDigital/core/test"
)

func reset() {
	os.Setenv("TMPENV", "HELLO WORLD")
	os.Unsetenv("TMPONE")
}

func ExampleTryLoadDev() {
	env.TryLoadDev()
}

func TestTryLoadDev(t *testing.T) {
	reset()
	env.TryLoadDev("testdata/one.env", "testdata/two.env")
	test.Equals(t, "HELLO WORLD", os.Getenv("TMPENV"))
	test.Equals(t, "ONE", os.Getenv("TMPONE"))
}

func ExampleTryOverloadDev() {
	env.TryOverloadDev()
}

func TestTryOverloadDev(t *testing.T) {
	reset()
	env.TryOverloadDev("testdata/one.env", "testdata/two.env")
	test.Equals(t, "TWO", os.Getenv("TMPENV"))
	test.Equals(t, "TWO", os.Getenv("TMPONE"))
}

func ExampleTryLoadTest() {
	env.TryLoadTest()
}

func TestTryLoadTest(t *testing.T) {
	reset()
	env.TryLoadTest("testdata/one.env", "testdata/two.env")
	test.Equals(t, "HELLO WORLD", os.Getenv("TMPENV"))
	test.Equals(t, "ONE", os.Getenv("TMPONE"))
}

func ExampleTryOverloadTest() {
	env.TryOverloadTest()
}

func TestTryOverloadTest(t *testing.T) {
	reset()
	env.TryOverloadTest("testdata/one.env", "testdata/two.env")
	test.Equals(t, "TWO", os.Getenv("TMPENV"))
	test.Equals(t, "TWO", os.Getenv("TMPONE"))
}
