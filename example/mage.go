// +build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "example"
	targets.ProjectType = "service"
	targets.DockerComposeDevDependencies = []string{"redis", "cockroach"}
	targets.DockerComposeTestDependencies = []string{"cockroach"}
	targets.DockerRunImage = targets.DockerRunImageMigrations
}
