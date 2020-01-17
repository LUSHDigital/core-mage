package targets

import (
	"context"
	"fmt"

	"github.com/magefile/mage/mg"
)

var (
	// MageTargetsRepo is the repository used for importing mage targets.
	MageTargetsRepo = "github.com/LUSHDigital/core-mage"
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
	if err := writeDotEnvFiles(); err != nil {
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
	if err := initGit(); err != nil {
		return err
	}
	return writeGitIgnoreFile()
}

// Gitlab sets up the gitlab pipeline
func (Setup) Gitlab(ctx context.Context) error {
	return writeGitlabCIFile()
}

// Dev is the namespace for actions related to the development environment.
type Dev mg.Namespace

// Start starts the development environment
func (Dev) Start(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "dev", DockerComposeDevFile)
	arg = append(arg, "up", "-d")
	arg = append(arg, DockerComposeDevDependencies...)
	return Exec(ComposeBin, arg...)
}

// Stop stops the development environment
func (Dev) Stop(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "dev", DockerComposeDevFile)
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

// Service starts the go service inside docker compose
func (Dev) Service(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "dev", DockerComposeDevFile)
	arg = append(arg, "up", "app")
	return Exec(ComposeBin, arg...)
}

// Build compiles the project inside a docker container
func Build(ctx context.Context) error {
	return Exec(DockerBin, "build", "-q", "--pull", ".")
}

// Tests is the namespace for actions related to the test environment.
type Tests mg.Namespace

// Run runs the project tests inside docker compose
func (Tests) Run(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "test", DockerComposeTestFile)
	arg = append(arg, "up")
	arg = append(arg,
		"--abort-on-container-exit",
		"--exit-code-from=app",
	)
	return Exec(ComposeBin, arg...)
}

// Reset sets the testing environment to its original state
func (Tests) Reset(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "test", DockerComposeTestFile)
	arg = append(arg, "down")
	return Exec(ComposeBin, arg...)
}

// Protos is the namespace for actions related to generating protobuffers.
type Protos mg.Namespace

// Add the protos submodule to the repository
func (Protos) Add(ctx context.Context) error {
	if err := addProtosSubmodule(); err != nil {
		return err
	}
	return nil
}

// Remove the protos submodule from the repository
func (Protos) Remove(ctx context.Context) error {
	if err := removeProtosSubmodule(); err != nil {
		return err
	}
	return nil
}

// Update the protos submodule
func (Protos) Update(ctx context.Context) error {
	if err := updateProtosSubmodule(); err != nil {
		return err
	}
	return nil
}

// Generate the protobuffers for this project
func (Protos) Generate(ctx context.Context) error {
	if err := genProtos(); err != nil {
		return err
	}
	return nil
}

// Mod is the namespace for actions related to installing modules.
type Mod mg.Namespace

// Core installs or upgrades all core packages
func (Mod) Core() error {
	libs := []string{
		"github.com/LUSHDigital/core-lush",
		"github.com/LUSHDigital/core",
		"github.com/LUSHDigital/spew",
	}
	return goget(libs...)
}

// Uuid installs or upgrades the uuid package
func (Mod) Uuid() error {
	libs := []string{
		"github.com/LUSHDigital/core",
		"github.com/LUSHDigital/spew",
	}
	return goget(libs...)
}

// Mysql installs or upgrades the mysql packages
func (Mod) Mysql() error {
	libs := []string{
		"github.com/LUSHDigital/core-sql",
		"github.com/LUSHDigital/scan",
		"github.com/go-sql-driver/mysql",
	}
	return goget(libs...)
}

// Postgres installs or upgrades the postgres packages
func (Mod) Postgres() error {
	libs := []string{
		"github.com/LUSHDigital/core-sql",
		"github.com/LUSHDigital/scan",
		"github.com/lib/pq",
	}
	return goget(libs...)
}

// Cockroach installs or upgrades the cockroach packages
func (Mod) Cockroach() error {
	libs := []string{
		"github.com/LUSHDigital/core-sql",
		"github.com/LUSHDigital/scan",
		"github.com/lib/pq",
	}
	return goget(libs...)
}

// Redis installs or upgrades the redis packages
func (Mod) Redis() error {
	libs := []string{
		"github.com/LUSHDigital/core-redis",
		"github.com/go-redis/redis",
		"github.com/alicebob/miniredis",
		"github.com/elliotchance/redismock",
	}
	return goget(libs...)
}

// Test runs the project tests inside docker compose
func Test(ctx context.Context) {
	mg.CtxDeps(ctx, Tests.Run)
}

// Install adds the dependencies into your vendor directory
func Install(ctx context.Context) error {
	args := []string{
		"run", "--rm",
		"-e", fmt.Sprintf("GOPROXY=%s", Environment.GoProxy()),
		"-v", "${PWD}:/repo",
		"-v", fmt.Sprintf("%s:/go/pkg/mod", Environment.GoModPath()),
		"-w", "/repo",
		DockerBuildImage,
	}
	return Exec(DockerBin, append(args,
		"go", "mod", "vendor",
	)...)
}

// Upgrade installs the latest version of the mage targets
func Upgrade(ctx context.Context) error {
	if err := Exec(GoBin, "get", "-u", MageTargetsRepo); err != nil {
		return err
	}
	return Exec(GoBin, "mod", "tidy")
}
