package env

import (
	"path/filepath"
	"testing"
)

var (
	// DefaultMigrationsPath is the default path to migrations relative from the project root.
	DefaultMigrationsPath = "service/database/migrations"
)

// MigrationsTestURL returns the migrations path relative to the project root.
func MigrationsTestURL(m *testing.M) string {
	return "file://" + filepath.Join(testRoot(), DefaultMigrationsPath)
}
