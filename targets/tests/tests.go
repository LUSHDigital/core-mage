package tests

import (
	"context"

	"github.com/LUSHDigital/core-mage/targets"
)

// Reset sets the testing environment to its original state
func Reset(ctx context.Context) error {
	arg := targets.BuildDockerComposeArgs(targets.ProjectName, targets.ProjectType, targets.DockerComposeTestFile)
	arg = append(arg, "down")
	return targets.Exec(targets.ComposeBin, arg...)
}
