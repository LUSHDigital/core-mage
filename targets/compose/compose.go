package compose

import "fmt"

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

	urlPattern string
}

// HostURL returns the fully qualified host url for the service
func (s Service) HostURL(host string) string {
	return fmt.Sprintf(s.urlPattern, host)
}

// ServiceManifest represent a collection of services for given nicknames.
type ServiceManifest map[string]Service

// EnvFor returns the environment key and value for a service on a given host.
func (sm ServiceManifest) EnvFor(host string, service string) (key string, value string) {
	return EnvVarNames[service], sm[service].HostURL(host)
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
	}
)
