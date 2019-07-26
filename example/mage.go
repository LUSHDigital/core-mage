// +build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "example"
	targets.ProjectType = "test"
	targets.DockerComposeDevDependencies = []string{"redis"}
	targets.DockerComposeTestDependencies = []string{"redis"}
}
