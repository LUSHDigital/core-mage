package targets

import (
	"os"
	"path"

	"github.com/magefile/mage/sh"
)

var (
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
)

const protosSubmoduleName = "protos"

func addProtosSubmodule() error {
	if err := addSubmodule(protosSubmoduleName, ProtoDefinitionsRepo, ProtoDefinitionsBranch); err != nil {
		return err
	}
	if err := initSubmodule(protosSubmoduleName); err != nil {
		return err
	}
	return updateSubmodule(protosSubmoduleName)
}

func removeProtosSubmodule() error {
	if err := deinitSubmodule(protosSubmoduleName); err != nil {
		return err
	}
	if err := removeSubmodule(protosSubmoduleName); err != nil {
		return err
	}
	return removeSubmoduleDir(protosSubmoduleName)
}

func updateProtosSubmodule() error {
	if err := initSubmodule(protosSubmoduleName); err != nil {
		return err
	}
	if err := updateSubmodule(protosSubmoduleName); err != nil {
		return err
	}
	return pullSubmodule(protosSubmoduleName, ProtoDefinitionsBranch)
}

func genProtos() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	output := path.Join(wd, ProtoOutputPath)
	protogenwd := path.Join(wd, submodulePath(protosSubmoduleName))

	if err := os.RemoveAll(path.Join(output, "protos")); err != nil {
		return err
	}
	for _, name := range ProtoServices {
		if err := genProtosFor(protogenwd, "service", name, output); err != nil {
			return err
		}
	}
	for _, name := range ProtoAggregators {
		if err := genProtosFor(protogenwd, "aggregators", name, output); err != nil {
			return err
		}
	}
	return nil
}

func genProtosFor(wd, namespace, name, output string) error {
	src := wd
	wd = path.Join(wd, "lush-protogen")
	if err := os.Chdir(wd); err != nil {
		return err
	}
	defer os.Chdir(Environment["PWD"])
	_, err := sh.Exec(
		Environment, os.Stdout, os.Stderr,
		GoBin, "run", "main.go", "--source", src, namespace, name, output,
	)
	return err
}
