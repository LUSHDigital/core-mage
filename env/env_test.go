package env_test

import (
	"os"
	"testing"

	"github.com/LUSHDigital/core-mage/env"
	"github.com/LUSHDigital/core/test"
)

func TestMain(m *testing.M) {
	reset()
	env.LoadTest(m, "env/testdata/test.load.env")
	env.OverloadTest(m, "env/testdata/test.overload.env")
	os.Exit(m.Run())
}

var m *testing.M

func ExampleLoadTest() {
	env.LoadTest(m, "infa/does-not-override.env")
}

func TestLoadTest(t *testing.T) {
	test.Equals(t, "loaded", os.Getenv("TMP_TEST_THREE"))
}

func ExampleOverloadTest() {
	env.OverloadTest(m, "infa/will-override.env")
}

func TestOverloadTest(t *testing.T) {
	test.Equals(t, "overloaded", os.Getenv("TMP_TEST_ONE"))
	test.Equals(t, "overloaded", os.Getenv("TMP_TEST_TWO"))
}

func ExampleLoad() {
	env.Load("infa/does-not-override.env")
}

func TestLoad(t *testing.T) {
	env.Load("testdata/local.load.env")
	test.Equals(t, "default", os.Getenv("TMP_LOCAL_ONE"))
	test.Equals(t, "default", os.Getenv("TMP_LOCAL_TWO"))
	test.Equals(t, "loaded", os.Getenv("TMP_LOCAL_THREE"))
}

func ExampleOverload() {
	env.Overload("infa/will-override.env")
}

func TestOverload(t *testing.T) {
	env.Overload("testdata/local.overload.env")
	test.Equals(t, "overloaded", os.Getenv("TMP_LOCAL_ONE"))
	test.Equals(t, "default", os.Getenv("TMP_LOCAL_TWO"))
	test.Equals(t, "loaded", os.Getenv("TMP_LOCAL_THREE"))
}

func reset() {
	os.Setenv("TMP_TEST_ONE", "default")
	os.Setenv("TMP_TEST_TWO", "default")
	os.Unsetenv("TMP_LOCAL_THREE")

	os.Setenv("TMP_LOCAL_ONE", "default")
	os.Setenv("TMP_LOCAL_TWO", "default")
	os.Unsetenv("TMP_TEST_THREE")
}
