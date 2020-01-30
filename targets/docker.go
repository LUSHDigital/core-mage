package targets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/LUSHDigital/core-mage/targets/compose"
	"gopkg.in/yaml.v2"
)

// DockerImage represents a reference to a remote docker image
type DockerImage = string

const (
	// DockerRunImageMigrations specifices the docker image for running a Go service with migrations.
	DockerRunImageMigrations DockerImage = "lushdigital/alpine-service:migrations"

	// DockerRunImageStandard specifices the docker image for running a Go regular service.
	DockerRunImageStandard DockerImage = "lushdigital/alpine-service:standard"

	// DockerBuildImageStandard specifices the docker image for building a regular Go service.
	DockerBuildImageStandard DockerImage = "lushdigital/alpine-golang:latest"
)

var (
	// DockerDir describes the path of the docker directory relative to the project root.
	DockerDir = "docker"

	// DockerBuildImage is the image used to build Go projects.
	DockerBuildImage = DockerBuildImageStandard

	// DockerRunImage is the image used to run Go projects.
	DockerRunImage = DockerRunImageStandard

	// DockerComposeTestFile configures the file that should be used for docker compose test environment.
	DockerComposeTestFile = path.Join(DockerDir, "test.yml")

	// DockerComposeTestEnvironment describes the environment variables that should be sent to docker compose apps in test.
	DockerComposeTestEnvironment = map[string]string{}

	// DockerComposeTestDependencies describe all dependencies that should be started in the docker compose test environment.
	DockerComposeTestDependencies = []string{}

	// DockerComposeDevFile configures the file that should be used for docker compose development environment.
	DockerComposeDevFile = path.Join(DockerDir, "dev.yml")

	// DockerComposeDevEnvironment describes the environment variables that should be sent to docker compose apps in development.
	DockerComposeDevEnvironment = map[string]string{}

	// DockerComposeDevDependencies describe all dependencies that should be started in the docker compose development environment.
	DockerComposeDevDependencies = []string{}
)

// BuildDockerComposeArgs will construct arguments for docker compose.
func BuildDockerComposeArgs(pname, ptype, namespace, file string) []string {
	return []string{
		"--no-ansi",
		"-p", fmt.Sprintf("%s_%s_%s", pname, ptype, namespace),
		"-f", file,
		"--project-directory", "${PWD}",
	}
}

func writeDockerfile() error {
	const dockerFileTmpl = "FROM %s\nFROM %s\n"
	raw := []byte(fmt.Sprintf(dockerFileTmpl, DockerBuildImage, DockerRunImage))
	return ioutil.WriteFile("Dockerfile", raw, 0664)
}

func writeDockerIgnorefile() error {
	var ignores = []string{
		"data/",
	}
	raw := []byte(strings.Join(ignores, "\n") + "\n")
	return ioutil.WriteFile(".dockerignore", raw, 0664)
}

func writeDockerDir() error {
	return os.MkdirAll(DockerDir, os.ModePerm)
}

func writeDockerComposeDev() error {
	var services = make(map[string]compose.Service)
	for _, name := range DockerComposeDevDependencies {
		if service, ok := compose.Services[name]; ok {
			services[name] = service
		}
	}
	app := compose.AppService
	app.Image = DockerBuildImageStandard
	app.DependsOn = DockerComposeDevDependencies
	services["app"] = app

	cf := compose.File{
		Version:  "2",
		Services: services,
	}
	raw, err := yaml.Marshal(&cf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(DockerComposeDevFile, raw, 0664)
}

func writeDockerComposeTest() error {
	var services = make(map[string]compose.Service)
	for _, name := range DockerComposeTestDependencies {
		if service, ok := compose.TestServices[name]; ok {
			services[name] = service
		}
	}

	app := compose.AppTestService
	app.Image = DockerBuildImage
	app.DependsOn = DockerComposeTestDependencies
	services["app"] = app

	cf := compose.File{
		Version:  "2",
		Services: services,
	}
	raw, err := yaml.Marshal(&cf)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(DockerComposeTestFile, raw, 0664)
}
