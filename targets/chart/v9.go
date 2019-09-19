package chart

import (
	"html/template"
	"io"
	"io/ioutil"

	"github.com/LUSHDigital/core-mage/targets/chart/templates"
	"github.com/Masterminds/sprig"
	"gopkg.in/yaml.v2"
)

// V9File represents the file structure of a standard chart values file.
type V9File struct {
	ChartVersion *string `yaml:"chartVersion"`

	ServiceLocation  *string `yaml:"serviceLocation"`
	ServiceNamespace *string `yaml:"serviceNamespace"`
	ServiceImage     *string `yaml:"serviceImage"`
	ServiceImageTag  *string `yaml:"serviceImageTag"`
	ServiceName      *string `yaml:"serviceName"`
	ServiceType      *string `yaml:"serviceType"`
	ServiceScope     *string `yaml:"serviceScope"`
	ServiceVersion   *string `yaml:"serviceVersion"`

	CIEnvironmentSlug     *string `yaml:"ciEnvironmentSlug"`
	CIPipelineID          *string `yaml:"ciPipelineId"`
	CIBuildID             *string `yaml:"ciBuildId"`
	CIEnvironmentHostname *string `yaml:"ciEnvironmentHostname"`

	LocalRedis  *int    `yaml:"localRedis"`
	MemoryStore *string `yaml:"memoryStore"`

	PrivateRegistryKey       *string                   `yaml:"privateRegistryKey"`
	Replicas                 *int                      `yaml:"replicas"`
	ProjectID                *string                   `yaml:"projectId"`
	PullPolicy               *string                   `yaml:"pullPolicy"`
	ProjectZone              *string                   `yaml:"projectZone"`
	DevelopmentVolumeMapping *int                      `yaml:"developmentVolumeMapping"`
	ServiceIngress           *int                      `yaml:"serviceIngress"`
	ServiceGrpcIngress       *int                      `yaml:"serviceGrpcIngress"`
	ServiceEnvironment       *string                   `yaml:"serviceEnvironment"`
	ClusterEnvironment       *string                   `yaml:"clusterEnvironment"`
	ServiceBranch            *string                   `yaml:"serviceBranch"`
	TLSKey                   *string                   `yaml:"tlsKey"`
	TLSCert                  *string                   `yaml:"tlsCert"`
	EnvVars                  map[string]string         `yaml:"envVars"`
	SecretEnvVars            map[string]V9SecretEnvVar `yaml:"secretEnvVars"`
	SetupJob                 *int                      `yaml:"setupJob"`
	SetupJobTimeout          *int                      `yaml:"setupJobTimeout"`
	JWTPublicURL             *string                   `yaml:"jwtPublicURL"`
	ExtraJobs                []V9ExtraJobs             `yaml:"extraJobs"`
	CockroachDB              *V9Cockroach              `yaml:"cockroachdb"`
}

// WriteTo provides a method for writing the V9File to a template.
func (f V9File) WriteTo(out io.Writer) (int64, error) {
	raw, err := templates.Box().MustString("v9.yaml")
	if err != nil {
		return 0, err
	}
	funcs := map[string]interface{}{
		"raw": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
	t := template.Must(template.New("v9.yaml").
		Funcs(sprig.FuncMap()).
		Funcs(funcs).
		Parse(raw))
	if err := t.ExecuteTemplate(out, "v9.yaml", &f); err != nil {
		return 0, err
	}
	return 0, nil
}

// V9Cockroach represents the configuration for cockroach in a values file.
type V9Cockroach struct {
	Enable       *bool `yaml:"enable"`
	Certificates *struct {
		Production *string `yaml:"production"`
		Staging    *string `yaml:"staging"`
	} `yaml:"certificates"`
}

// V9ExtraJobs represents extra jobs for the chart.
type V9ExtraJobs struct {
	Name     *string           `yaml:"name"`
	Command  []string          `yaml:"command"`
	EnvVars  map[string]string `yaml:"envVars"`
	Schedule *string           `yaml:"schedule"`
	Image    *struct {
		Repository string `yaml:"repository"`
		Tag        string `yaml:"tag"`
	} `yaml:"certificates"`
	EnableDB      *bool   `yaml:"enableDB"`
	RestartPolicy *string `yaml:"restartPolicy"`
}

// V9SecretEnvVar represents the key for environment variables derived from secrets.
type V9SecretEnvVar struct {
	Name       *string `yaml:"name"`
	SecretName *string `yaml:"secretName"`
	Key        *string `yaml:"key"`
}

// ReadV9File loads a version file from disk.
func ReadV9File(filename string) (V9File, error) {
	var vf V9File
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return vf, err
	}
	if err := yaml.Unmarshal(raw, &vf); err != nil {
		return vf, err
	}
	return vf, nil
}
