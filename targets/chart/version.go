package chart

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Version represents the value type for the chart version enum.
type Version int

const (
	// V01 is the previous chart version for developing on the backend-specific infrastructure.
	V01 Version = iota
	// V9 is the latest chart version for developing on the backend-specific infrastructure.
	V9
	// VUnknown is an unknown chart version type.
	VUnknown
)

// ReadVersionFile loads a version file from disk.
func ReadVersionFile(filename string) (VersionFile, error) {
	var vf VersionFile
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return vf, err
	}
	if err := yaml.Unmarshal(raw, &vf); err != nil {
		return vf, err
	}
	return vf, nil
}

// VersionFile represents the file structure of a chart file only with version.
type VersionFile struct {
	ChartVersion string `yaml:"chartVersion"`
}

// Version returns the correct version to be using.
func (f VersionFile) Version() Version {
	return f.parseVersion(f.ChartVersion)
}

func (f VersionFile) parseVersion(version string) Version {
	switch version {
	case "0.1.0":
		return V01
	case "9.0.0-stable", "9.0.0":
		return V9
	}
	return VUnknown
}
