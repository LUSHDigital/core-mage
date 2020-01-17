// Package env is used to load environment variables in the context of using the core-mage targets and development environment.
// It provides convenient functions for loading or overloading.
package env

import (
	"log"

	"github.com/joho/godotenv"
)

var (
	// DevFiles represents the development files to be loaded/overloaded in order.
	DevFiles = []string{
		"infra/personal.env",
		"infra/common.env",
		"infra/local.dev.env",
	}

	// TestFiles represents the test files to be loaded/overloaded in order.
	TestFiles = []string{
		"infra/personal.env",
		"infra/common.env",
		"infra/local.test.env",
	}
)

// TryLoadDev will attempt to load the development environment variables.
// This WILL NOT OVERRIDE a variable that already exists and those set prior to this call will remain.
// For variables set in multiple files passed to this function, the FIRST one will prevail.
func TryLoadDev(paths ...string) {
	load(append(DevFiles, paths...)...)
}

// TryOverloadDev will attempt to load the development environment variables.
// This WILL OVERRIDE a variable that already exists and those set prior to this call will be replaced if set here.
// For env variables set in multiple files passed to this function, the LAST one will prevail.
func TryOverloadDev(paths ...string) {
	overload(append(DevFiles, paths...)...)
}

// TryLoadTest will attempt to load the test environment variables.
// This WILL NOT OVERRIDE a variable that already exists and those set prior to this call will remain.
// For variables set in multiple files passed to this function, the FIRST one will prevail.
func TryLoadTest(paths ...string) {
	load(append(DevFiles, paths...)...)
}

// TryOverloadTest will attempt to load the test environment variables.
// This WILL OVERRIDE a variable that already exists and those set prior to this call will be replaced if set here.
// For env variables set in multiple files passed to this function, the LAST one will prevail.
func TryOverloadTest(paths ...string) {
	overload(append(DevFiles, paths...)...)
}

func load(paths ...string) {
	for _, p := range paths {
		if err := godotenv.Load(p); err != nil {
			log.Printf("could not load environment file: %s: skipping...\n", p)
		}
	}
}

func overload(paths ...string) {
	for _, p := range paths {
		if err := godotenv.Overload(p); err != nil {
			log.Printf("could not overload environment file: %s: skipping...\n", p)
		}
	}
}
