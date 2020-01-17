package targets

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/magefile/mage/sh"

	"github.com/LUSHDigital/core-mage/targets/chart"
	"github.com/LUSHDigital/core-mage/targets/compose"

	"github.com/joho/godotenv"
)

var (
	// InfraDir represents the directory used for infrastructure files.
	InfraDir = "infra"

	// GCPStageChartFile configures the file path to the gcp chart for stage.
	GCPStageChartFile = path.Join(InfraDir, "stage.gcp.yaml")

	// GCPProdChartFile configures the file path to the gcp chart for prod.
	GCPProdChartFile = path.Join(InfraDir, "prod.gcp.yaml")
)

type chartData struct {
	Name string
	Type string
}

func writeInfraDir() error {
	return os.MkdirAll(InfraDir, os.ModePerm)
}

func writeStageChart() error {
	return writeChart(GCPStageChartFile)
}

func writeProdChart() error {
	return writeChart(GCPProdChartFile)
}

func writeDotEnvFiles() error {
	if err := MoveLegacyDotEnvFiles(); err != nil {
		return err
	}
	if err := WriteDotEnvCommon(); err != nil {
		return err
	}
	if err := WriteDotEnvPersonal(); err != nil {
		return err
	}
	if err := WriteDotEnvLocalDev(); err != nil {
		return err
	}
	if err := WriteDotEnvLocalTest(); err != nil {
		return err
	}
	if err := WriteDotEnvComposeDev(); err != nil {
		return err
	}
	if err := WriteDotEnvComposeTest(); err != nil {
		return err
	}
	return nil
}

// MoveLegacyDotEnvFiles will take old dot env files and move them to their newer counterparts location
func MoveLegacyDotEnvFiles() error {
	mv := func(src, dst string) {
		if err := sh.Copy(dst, src); err != nil {
			return
		}
		if err := sh.Rm(src); err != nil {
			fmt.Printf("cannot remove file %q: %v\n", src, err)
			return
		}
		fmt.Printf("moved file %q to %q\n", src, dst)
	}
	mv("infra/.env", "infra/local.dev.env")
	mv("infra/local.env", "infra/personal.env")
	mv("infra/dev.env", "infra/compose.dev.env")
	mv("infra/test.env", "infra/compose.test.env")
	return nil
}

// WriteDotEnvCommon will write the configuration file for the local development and test environment.
func WriteDotEnvCommon() error {
	var vars = make(map[string]string)
	return WriteEnvFile(
		path.Join(InfraDir, "common.env"),
		vars,
		"Environment variables for the development and test environments both locally and in docker compose",
	)
}

// WriteDotEnvPersonal will write the configuration file for the local development and test environment and within docker compose.
func WriteDotEnvPersonal() error {
	var vars = make(map[string]string)
	return WriteEnvFile(
		path.Join(InfraDir, "personal.env"),
		vars,
		"Environment variables for the development and test environments both locally and in docker compose",
		"This file should be IGNORED by the SCM",
	)
}

// WriteDotEnvLocalDev will write the configuration file for the local development environment.
func WriteDotEnvLocalDev() error {
	var vars = make(map[string]string)
	for _, dep := range DockerComposeDevDependencies {
		k, v := compose.Services.EnvFor(LocalHost, dep)
		vars[k] = v
	}
	if DockerRunImage == DockerRunImageMigrations {
		vars[MigrationsURLEnvVar] = MigrationsURLLocal
	}
	return WriteEnvFile(
		path.Join(InfraDir, "local.dev.env"),
		vars,
		"Environment variables for the development environment locally",
	)
}

// WriteDotEnvLocalTest will write the configuration file for the local test environment.
func WriteDotEnvLocalTest() error {
	var vars = make(map[string]string)
	for _, dep := range DockerComposeTestDependencies {
		k, v := compose.Services.EnvFor(LocalHost, dep)
		vars[k] = v
	}
	if DockerRunImage == DockerRunImageMigrations {
		vars[MigrationsURLEnvVar] = MigrationsURLLocal
	}
	return WriteEnvFile(
		path.Join(InfraDir, "local.test.env"),
		vars,
		"Environment variables for the test environment locally",
	)
}

// WriteDotEnvComposeDev will write the configuration file for the development environment within docker compose.
func WriteDotEnvComposeDev() error {
	var vars = make(map[string]string)
	for _, dep := range DockerComposeDevDependencies {
		k, v := compose.Services.EnvFor(dep, dep)
		vars[k] = v
	}
	if DockerRunImage == DockerRunImageMigrations {
		vars[MigrationsURLEnvVar] = MigrationsURLDev
	}
	return WriteEnvFile(
		path.Join(InfraDir, "compose.dev.env"),
		vars,
		"Environment variables for the development environment in docker compose",
	)
}

// WriteDotEnvComposeTest will write the configuration file for the test environment within docker compose.
func WriteDotEnvComposeTest() error {
	var vars = make(map[string]string)
	for _, dep := range DockerComposeTestDependencies {
		k, v := compose.Services.EnvFor(dep, dep)
		vars[k] = v
	}
	if DockerRunImage == DockerRunImageMigrations {
		vars[MigrationsURLEnvVar] = MigrationsURLTest
	}
	return WriteEnvFile(
		path.Join(InfraDir, "compose.test.env"),
		vars,
		"Environment variables for the test environment in docker compose",
	)
}

// WriteEnvFile writes an environment file to disk but retain manual changes that have been made.
func WriteEnvFile(filename string, vars map[string]string, comments ...string) error {
	filevars, err := godotenv.Read(filename)
	if err != nil {
		fmt.Printf("env file %q does not exist: creating...\n", filename)
	}
	var allvars = make(map[string]string)
	for k, v := range vars {
		allvars[strings.ToUpper(k)] = v
	}
	for k, v := range filevars {
		allvars[strings.ToUpper(k)] = v
	}
	buf := bytes.NewBuffer(nil)
	for _, comment := range comments {
		fmt.Fprintf(buf, "# %s\n", comment)
	}
	for k, v := range allvars {
		fmt.Fprintf(buf, "%s=%s\n", k, v)
	}
	return ioutil.WriteFile(filename, buf.Bytes(), 0664)
}

func writeChart(filename string) error {
	_, err := os.Stat(filename)
	if err != nil && os.IsExist(err) {
		return err
	}
	buf := bytes.NewBuffer(nil)
	if os.IsNotExist(err) {
		chartVersion := "9.0.0-stable"
		pullPolicy := "Always"
		replicas := 3
		v9f := chart.V9File{
			ChartVersion: &chartVersion,
			ServiceName:  &ProjectName,
			ServiceType:  &ProjectType,
			Replicas:     &replicas,
			PullPolicy:   &pullPolicy,
		}
		if _, err := v9f.WriteTo(buf); err != nil {
			return err
		}
	} else {
		vf, err := chart.ReadVersionFile(filename)
		if err != nil {
			return err
		}
		switch vf.Version() {
		case chart.V9:
			v9f, err := chart.ReadV9File(filename)
			if err != nil {
				return err
			}
			_, err = v9f.WriteTo(buf)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown chart version: %q", vf.ChartVersion)
		}
	}
	return ioutil.WriteFile(filename, []byte(strings.TrimSpace(buf.String())), 0664)
}
