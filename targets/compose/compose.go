package compose

import (
	"fmt"
	"strings"
)

// File represents the file structure of a docker compose file.
type File struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services,omitempty"`
}

// Build represents the structure of a build directive in a docker compose file.
type Build struct {
	Context    string `yaml:"context,omitempty"`
	Dockerfile string `yaml:"dockerfile,omitempty"`
}

// Service represents the structure of a docker compose service.
type Service struct {
	Image      string `yaml:"image,omitempty"`
	Command    string `yaml:"command,omitempty"`
	Build      Build  `yaml:"build,omitempty"`
	WorkingDir string `yaml:"working_dir,omitempty"`
	Restart    string `yaml:"restart,omitempty"`

	Logging map[string]string `yaml:"logging,omitempty"`

	Volumes []string `yaml:"volumes,omitempty"`
	Ports   []string `yaml:"ports,omitempty"`

	EnvFile     string            `yaml:"env_file,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`

	DependsOn []string `yaml:"depends_on,omitempty"`

	ExternalURLPattern string `yaml:"-"`
	InternalURLPattern string `yaml:"-"`
}

// HostURL returns the fully qualified host url for the service
func (s Service) HostURL(host string) string {
	pattern := s.ExternalURLPattern
	if pattern == "" {
		pattern = "%s:0"
	}
	return fmt.Sprintf(s.ExternalURLPattern, host)
}

// InternalHostURL returns the fully qualified host url for the service
func (s Service) InternalHostURL(host string) string {
	pattern := s.InternalURLPattern
	if pattern == "" {
		pattern = "%s:0"
	}
	return fmt.Sprintf(s.InternalURLPattern, host)
}

// ServiceManifest represent a collection of services for given nicknames.
type ServiceManifest map[string]Service

// EnvFor returns the environment key and value for a service on a given host.
func (sm ServiceManifest) EnvFor(host string, service string) (key string, value string) {
	key, ok := EnvVarNames[service]
	if !ok {
		key = fmt.Sprintf(strings.ToUpper("%s_URL"), service)
	}
	return key, sm[service].HostURL(host)
}

// InternalEnvFor returns the environment key and value for a service on a given host.
func (sm ServiceManifest) InternalEnvFor(host string, service string) (key string, value string) {
	key, ok := EnvVarNames[service]
	if !ok {
		key = fmt.Sprintf(strings.ToUpper("%s_URL"), service)
	}
	return key, sm[service].InternalHostURL(host)
}

var (
	// Services represents all available docker compose services for the development environment.
	Services = ServiceManifest{
		"pg":          PostgresService,
		"postgres":    PostgresService,
		"postgresql":  PostgresService,
		"mysql":       MySQLService,
		"cockroach":   CockroachService,
		"cockroachdb": CockroachService,
		"redis":       RedisService,
		"mongo":       MongoService,
		"mongodb":     MongoService,
		"minio":       MinioService,
		"rabbitmq":    RabbitMQService,
	}

	// TestServices represents all available docker compose services for the testing environment.
	TestServices = ServiceManifest{
		"pg":          PostgresTestService,
		"postgres":    PostgresTestService,
		"postgresql":  PostgresTestService,
		"mysql":       MySQLTestService,
		"cockroach":   CockroachTestService,
		"cockroachdb": CockroachTestService,
		"redis":       RedisTestService,
		"mongo":       MongoTestService,
		"mongodb":     MongoTestService,
		"minio":       MinioTestService,
	}

	// EnvVarNames represents all available env var names for the given service.
	EnvVarNames = map[string]string{
		"pg":          "POSTGRES_URL",
		"postgres":    "POSTGRES_URL",
		"postgresql":  "POSTGRES_URL",
		"mysql":       "MYSQL_URL",
		"cockroach":   "COCKROACH_URL",
		"cockroachdb": "COCKROACH_URL",
		"redis":       "REDIS_URL",
		"mongo":       "MONGO_URL",
		"mongodb":     "MONGO_URL",
		"minio":       "MINIO_URL",
		"rabbitmq":    "RABBITMQ_URL",
	}
)
