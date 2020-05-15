// +build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "example"
	targets.ProjectType = "service"
	targets.DockerComposeDevDependencies = []string{"redis", "cockroach", "minio", "rabbit"}
	targets.DockerComposeTestDependencies = []string{"rabbit", "cockroach"}
	targets.DockerRunImage = targets.DockerRunImageMigrations
	targets.ProtoServices = []string{"products"}
	targets.ProtoDefinitionsBranch = "master"

	// Used to account for the fact that we're importing a dependency from the parent package.
	targets.InstallVolume = "${PWD}/..:/repo"
	targets.InstallWorkDir = "/repo/example"
}
