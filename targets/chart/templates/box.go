// Package templates provide templates for generating helm values files.
//go:generate packr
package templates

import "github.com/gobuffalo/packr"

// Box returns the packr box.
func Box() packr.Box {
	return packr.NewBox(".")
}
