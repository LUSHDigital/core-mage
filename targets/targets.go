package targets

import (
	"context"
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

// Setup is the namespace for actions related to setting up the project.
type Setup mg.Namespace

// All performs all the build steps
func (s Setup) All(ctx context.Context) {
	mg.CtxDeps(ctx,
		s.Docker,
		s.Gitlab,
		s.Infra,
		s.Git,
	)
}

// Infra installs the infrastructure dependencies
func (Setup) Infra(ctx context.Context) error {
	if err := writeInfraDir(); err != nil {
		return err
	}
	if err := writeStageChart(); err != nil {
		return err
	}
	if err := writeProdChart(); err != nil {
		return err
	}
	if err := writeDotEnv(); err != nil {
		return err
	}
	if err := writeDotEnvDev(); err != nil {
		return err
	}
	if err := writeDotEnvTest(); err != nil {
		return err
	}
	if err := writeDotEnvLocal(); err != nil {
		return err
	}
	return nil
}

// Docker installs the docker dependencies
func (Setup) Docker(ctx context.Context) error {
	if err := writeDockerfile(); err != nil {
		return err
	}
	if err := writeDockerIgnorefile(); err != nil {
		return err
	}
	if err := writeDockerDir(); err != nil {
		return err
	}
	if err := writeDockerComposeDev(); err != nil {
		return err
	}
	if err := writeDockerComposeTest(); err != nil {
		return err
	}
	return nil
}

// Git sets up git inside the project
func (Setup) Git(ctx context.Context) error {
	return writeGitIgnoreFile()
}

// Gitlab sets up the gitlab pipeline
func (Setup) Gitlab(ctx context.Context) error {
	if err := Exec(GitBin, "init"); err != nil {
		return err
	}
	return writeGitlabCIFile()
}

// Dev is the namespace for actions related to the development environment.
type Dev mg.Namespace

// Start starts the development environment
func (Dev) Start(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, DockerComposeDevFile)
	arg = append(arg, "up", "-d")
	arg = append(arg, DockerComposeDevDependencies...)
	return Exec(ComposeBin, arg...)
}

// Stop stops the development environment
func (Dev) Stop(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, DockerComposeDevFile)
	arg = append(arg, "stop")
	arg = append(arg, DockerComposeDevDependencies...)
	return Exec(ComposeBin, arg...)
}

// Restart will first stop then start the development environment
func (Dev) Restart(ctx context.Context) {
	mg.SerialCtxDeps(ctx,
		Dev.Stop,
		Dev.Start,
	)
}

// Service starts the go service
func (Dev) Service(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, DockerComposeDevFile)
	arg = append(arg, "up", "app")
	return Exec(ComposeBin, arg...)
}

// Build compiles the project inside a docker container
func Build(ctx context.Context) error {
	return Exec(DockerBin, "build", "-q", "--pull", ".")
}

// Install runs the tests for the project
func Install(ctx context.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	args := []string{
		"run", "--rm",
		"-e", fmt.Sprintf("GOPROXY=%q", Environment["GOPROXY"]),
		"-v", fmt.Sprintf("%s:/repo", wd),
		"-v", fmt.Sprintf("%s:/go/pkg/mod", Environment.GoModPath()),
		"-w", "/repo",
		DockerBuildImage,
	}
	return Exec(DockerBin, append(args, "go", "mod", "vendor")...)
}

// Test runs the tests for the project
func Test(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, DockerComposeTestFile)
	arg = append(arg, "up")
	arg = append(arg,
		"--abort-on-container-exit",
		"--exit-code-from app",
	)
	return Exec(ComposeBin, arg...)
}
