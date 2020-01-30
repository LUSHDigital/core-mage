// Package env is used to load environment variables in the context of using the core-mage targets and development environment.
// It provides convenient functions for loading or overloading.
package env

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/LUSHDigital/core-mage/env/internal/mod"

	"github.com/joho/godotenv"
)

var (
	// DevFiles represents the development files to be loaded in order.
	DevFiles = []string{
		"infra/personal.env",
		"infra/local.dev.env",
		"infra/common.env",
		".env",
	}

	// TestFiles represents the test files to be loaded in order.
	TestFiles = []string{
		"infra/local.test.env",
		"infra/common.env",
	}
)

// LoadDefault will attempt to load the default environment.
func LoadDefault() {
	for _, fpath := range DevFiles {
		if err := godotenv.Load(fpath); err != nil {
			log.Printf("could not load environment file: %s: skipping...\n", fpath)
		}
	}
}

// Load tries to load given env files, leaving current environment variables intact.
func Load(paths ...string) {
	for _, fpath := range paths {
		if err := godotenv.Load(fpath); err != nil {
			log.Printf("could not load environment file: %s: skipping...\n", fpath)
		}
	}
}

// Overload tries to overload given env files, overriding current environment variables.
func Overload(paths ...string) {
	for _, fpath := range paths {
		if err := godotenv.Overload(fpath); err != nil {
			log.Printf("could not overload environment file: %s: skipping...\n", fpath)
		}
	}
}

// MustLoadDefaultTest will attempt to load the default test environment.
func MustLoadDefaultTest(m *testing.M) {
	dir := testRoot()
	for _, fpath := range TestFiles {
		fullpath := filepath.Join(dir, fpath)
		if err := godotenv.Load(fullpath); err != nil {
			log.Fatalf("could not load environment file: %s: %v", fullpath, err)
		}
	}
}

// LoadTest tries to load given env files, leaving current environment variables intact.
func LoadTest(m *testing.M, paths ...string) {
	dir := testRoot()
	for _, fpath := range paths {
		fullpath := filepath.Join(dir, fpath)
		if err := godotenv.Load(fullpath); err != nil {
			log.Printf("could not load environment file: %s: skipping...\n", fullpath)
		}
	}
}

// OverloadTest tries to overload given env files, overriding current environment variables.
func OverloadTest(m *testing.M, paths ...string) {
	dir := testRoot()
	for _, fpath := range paths {
		fullpath := filepath.Join(dir, fpath)
		if err := godotenv.Overload(fullpath); err != nil {
			log.Printf("could not overload environment file: %s: skipping...\n", fullpath)
		}
	}
}

func testRoot() string {
	dir, err := mod.ModulePath()
	if err != nil {
		log.Fatalf("cannot find module path: %v", err)
	}
	return dir
}
