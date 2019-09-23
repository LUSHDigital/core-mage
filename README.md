# Magefile targets
A set of Magefile targets for managing services on the LUSH platform.

## What is mage?
> Mage is a make/rake-like build tool using Go. You write plain-old go functions, and Mage automatically uses them as Makefile-like runnable targets.
>
> – [Nate Finch](https://github.com/natefinch) creator of [Mage](https://magefile.org/)

## Installation
Start by installing the mage binary [follwing the documentation on their website](https://magefile.org/). The next step would be to install your mage targets inside a project with go modules enabled.

```bash
$ go get -u github.com/LUSHDigital/core-mage@latest
```

Once you've installed mage and the core-mage library, you can create a `mage.go` in the root of your project. Don't forget to set your project variables in the init function.

```go
// +build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "my-awesome-app"
	targets.ProjectType = "service"
}
```

## Usage

### Setup
To setup a project using mage, you can simply run the setup command in the root of your repository. This will do things like initialising your git repository, creating your helm values files, creating docker compose files for testing and development and a few other things you need to get going with a LUSH platform enabled service.

```bash
$ mage setup:all
```

#### Reference: infra
Once the setup has completed, you will notice under the `infra/` folder that some files have been created. Here is a quick reference:

- `local.env` defines environment variables when running your application with `go run` or `go build`
- `dev.env` defines environment variables when running your application with `mage dev:service`
- `test.env` defines environment variables when running your application with `mage test`
- `staging.gcp.yml` defines environment variables & configuration options when running your application within the `staging` cluster
- `prod.gcp.yml` defines environment variables & configuration options when running your application within the `production` cluster

Note that `dev.env` and `test.env` are useful to avoid having to fiddle with docker-compose definitions.

#### Reference: protos

Note that the following configuration options are available when dealing with the `mage protos` namespace.

```
// ProtoDefinitionsRepo is the central repository used for proto definitions.
ProtoDefinitionsRepo = "git@gitlab.com:LUSHDigital/soa/models/rpc.git"

// ProtoDefinitionsBranch is the branch of the protos repository to check out.
ProtoDefinitionsBranch = "master"

// ProtoOutputPath is the path where the generated protos should be output to.
ProtoOutputPath = "service"

// ProtoServices are the service protobuffers that should be generated with lush-protogen.
ProtoServices = []string{}

// ProtoAggregators are the aggregator protobuffers that should be generated with lush-protogen.
ProtoAggregators = []string{}
```

`ProtoDefinitionsBranch` can be used to point to a development branch, rather than the default (master).
This allows working with volatile definitions which aren't ready for general use.


### Test
Every project should have tests, yours is no exception. These mage targets make it easy to both run and manage your test environment. To run your tests, simply run the target and let mage to the rest.

```
$ mage test
```

If you need to add dependencies, you can add them into the init function inside your magefile and re-run your setup targets which will add them to your docker compose file.

```go
func init() {
    ...
    targets.DockerComposeTestDependencies = []string{"mysql"}
}
```

### Develop
These mage targets will not only manage your testing environment, but your development environment as-well. If you want to to add development dependencies, add them to your init function inside your magefile and re-run the setup target.

```go
func init() {
    ...
    targets.DockerComposeDevDependencies = []string{"redis", "mysql"}
}
```

To only start your dependencies without running your application run the start target. This spins up your development dependencies in the background and you'll be able to connect to them from your machine. Mage manages and populates dotenv files in your `infra/` directory with the appropriate connection strings to your dependencies.

```bash
$ mage dev:start
```

You can also run your service inside of docker compose and have it run in the foreground by running the service target. The service will be killed if you cancel the interupt the process and you'll be able to see your application logs in your terminal. Very good to quickly get your application running without much faff.

```bash
$ mage dev:service
```

## Upgrading
We've provided simple tooling for having mage be self-upgrading. Run the target and mage will take care of the rest.

```bash
$ mage upgrade
```

## Contributing
If you're thinking about contributing to this repository, first of all, thank you! Second of all is to tell you to clone the repository outside of your `GOPATH` since this is a module based project. Once you're set-up, you can make the changes you want and try your targets in the `example/` directory, which is set up to point at your local targets package.

**Happy wizardry!**
