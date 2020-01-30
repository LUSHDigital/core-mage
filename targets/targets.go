package targets

import (
	"context"
	"fmt"

	"github.com/magefile/mage/mg"
)

var (
	// MageTargetsRepo is the repository used for importing mage targets.
	MageTargetsRepo = "github.com/LUSHDigital/core-mage"

	// InstallVolume is used to set what volume to use during the `mage install` target.
	InstallVolume string = "${PWD}:/repo"

	// InstallWorkDir is used to set what work directory to use during the `mage install` target.
	InstallWorkDir string = "/repo"
)

// Setup performs setup of the project according to your magefile configuration.
func Setup(ctx context.Context) error {
	if err := setupDocker(); err != nil {
		return err
	}
	if err := setupGitlab(); err != nil {
		return err
	}
	if err := setupInfra(); err != nil {
		return err
	}
	if err := setupGit(); err != nil {
		return err
	}
	return nil
}

// Dev is the namespace for actions related to the development environment.
type Dev mg.Namespace

// Start starts the development environment in docker compose
func (Dev) Start(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "dev", DockerComposeDevFile)
	arg = append(arg, "up", "-d")
	arg = append(arg, DockerComposeDevDependencies...)
	return Exec(ComposeBin, arg...)
}

// Stop stops the development environment in docker compose
func (Dev) Stop(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "dev", DockerComposeDevFile)
	arg = append(arg, "stop")
	arg = append(arg, DockerComposeDevDependencies...)
	return Exec(ComposeBin, arg...)
}

// Restart restarts the development environment in docker compose
func (Dev) Restart(ctx context.Context) {
	mg.SerialCtxDeps(ctx,
		Dev.Stop,
		Dev.Start,
	)
}

// Run runs the service inside docker compose
func (Dev) Run(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "dev", DockerComposeDevFile)
	arg = append(arg, "run")
	arg = append(arg,
		"--service-ports",
		"--use-aliases",
	)
	arg = append(arg, "app", "go", "run", "-mod=vendor", "./service")
	return Exec(ComposeBin, arg...)
}

// Build compiles the project inside a docker container
func Build(ctx context.Context) error {
	return Exec(DockerBin, "build", "-q", "--pull", ".")
}

// Tests is the namespace for actions related to the test environment.
type Tests mg.Namespace

// Start starts the test environment
func (Tests) Start(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "test", DockerComposeTestFile)
	arg = append(arg, "up", "-d")
	arg = append(arg, DockerComposeTestDependencies...)
	return Exec(ComposeBin, arg...)
}

// Reset returns the testing environment to its orginal state
func (Tests) Reset(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "test", DockerComposeTestFile)
	arg = append(arg, "down")
	return Exec(ComposeBin, arg...)
}

// Run runs the project tests inside docker compose
func (Tests) Run(ctx context.Context) error {
	arg := BuildDockerComposeArgs(ProjectName, ProjectType, "test", DockerComposeTestFile)
	arg = append(arg, "run")
	arg = append(arg,
		"--service-ports",
		"--use-aliases",
	)
	arg = append(arg, "app", "go", "test", "-mod=vendor", "-v", "-cover", "./...")
	return Exec(ComposeBin, arg...)
}

// Protos is the namespace for actions related to generating protobuffers.
type Protos mg.Namespace

// Add adds the protos submodule to the repository
func (Protos) Add(ctx context.Context) error {
	if err := addProtosSubmodule(); err != nil {
		return err
	}
	return nil
}

// Remove removes the protos submodule from the repository
func (Protos) Remove(ctx context.Context) error {
	if err := removeProtosSubmodule(); err != nil {
		return err
	}
	return nil
}

// Update updates the protos submodule in the repository
func (Protos) Update(ctx context.Context) error {
	if err := updateProtosSubmodule(); err != nil {
		return err
	}
	mg.CtxDeps(ctx, Protos.Generate)
	return nil
}

// Generate generates codes from the protos submodule into the repository
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

// Install installs the go module dependencies in the vendor directory
func Install(ctx context.Context) error {
	args := []string{
		"run", "--rm",
	}
	if proxy := Environment.GoProxy(); proxy != "" {
		args = append(args, "-e", fmt.Sprintf("GOPROXY=%s", proxy))
	}
	if modpath := Environment.GoModPath(); modpath != "" {
		args = append(args, "-v", fmt.Sprintf("%s:/go/pkg/mod", modpath))
	}
	args = append(args,
		"-v", InstallVolume,
		"-w", InstallWorkDir,
		DockerBuildImage,
	)
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
