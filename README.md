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

- `common.env` contain variables for the development and environments both locally and in docker compose
- `personal.env` personal contain variables for the development and environments both locally and in docker compose and is ignored by git by default
- `compose.dev.env` contain variables for the development environment within docker compose when running `mage dev:service`
- `compose.test.env` contain variables for the test environment within docker compose when running `mage test`
- `staging.gcp.yml` contain variables & configuration options when running your application within the `staging` cluster
- `prod.gcp.yml` contain variables & configuration options when running your application within the `production` cluster

Note that `compose.dev.env` and `compose.test.env` are useful to avoid having to fiddle with the docker compose definitions.

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

### Running tests on your local machine
Sometimes you want to be able to run partial tests, add test flags or have your IDE run the tests for you. To achieve this you can start your test dependencies by running the `tests:start` target. This will keep them running in the background and expose their ports to the host machine.

```
$ mage tests:start
```

### Resetting your test environment to its original state
If your tests have gotten your database in a broken state and you don't know why, you can always reset the entire test environment using the `tests:reset` target to bring it down. Remember you need to run `tests:start` again if you want to run your tests on the host machine.

```
$ mage tests:reset
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

### Protos
Mage depends on two commandline utilities to generate the protobuffer files, `protobuf` and `protoc-gen-go`, so ensure these are installed and updated before generating protobuffer files. Protobuffer files from a git respository can be included in your project by first setting up your mage.go file to include them. 

```go
//+build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "products"
	targets.ProjectType = "service"
	targets.ProtoServices = []string{"products"}
}
```

The first time you generate protobuffers to your project you need to run `mage protos:add` and every time you need to do an update, run `mage protos:update` which will pull the latest version of this repository and re-generate all protos into go code within your project.

```
mage protos:update
lush-protogen generating go protos into "protos/service/products" package
	- ./modules/protos/protos/service/products/classifications.proto
	- ./modules/protos/protos/service/products/enums.proto
	- ./modules/protos/protos/service/products/images.proto
	- ./modules/protos/protos/service/products/indices.proto
	- ./modules/protos/protos/service/products/ingredients.proto
	- ./modules/protos/protos/service/products/items.proto
	- ./modules/protos/protos/service/products/jurisdictions.proto
	- ./modules/protos/protos/service/products/markets_service.proto
	- ./modules/protos/protos/service/products/products.proto
	- ./modules/protos/protos/service/products/products_service.proto
	- ./modules/protos/protos/service/products/types.proto
writing files to ./service/protos/service/products
```

Protobuffer files can also be generated directly by running `mage protos:generate`.

#### Branches
By default the protobuffer files are generated for the master branch of your git repository. To target a different branch you need to define the `targets.ProtoDefinitionsBranch` in your mage.go file. Then update and generate your protobuffers as above.

```go
//+build mage

package main

import (
	// mage:import
	"github.com/LUSHDigital/core-mage/targets"
)

func init() {
	targets.ProjectName = "products"
	targets.ProjectType = "service"
	targets.ProtoServices = []string{"products"}
	targets.ProtoDefinitionsBranch = "feat/tag-groups"
}
```


## Upgrading
We've provided simple tooling for having mage be self-upgrading. Run the target and mage will take care of the rest.

```bash
$ mage upgrade
```

## Contributing
If you're thinking about contributing to this repository, first of all, thank you! Second of all is to tell you to clone the repository outside of your `GOPATH` since this is a module based project. Once you're set-up, you can make the changes you want and try your targets in the `example/` directory, which is set up to point at your local targets package.

**Happy wizardry!**
