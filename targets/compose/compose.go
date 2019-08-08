package compose

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
}

var (
	// Services represents all available docker compose services for the development environment.
	Services = map[string]Service{
		"pg":          PostgresService,
		"postgres":    PostgresService,
		"postgresql":  PostgresService,
		"mysql":       MySQLService,
		"cockroach":   CockroachService,
		"cockroachdb": CockroachService,
		"redis":       RedisService,
	}

	// TestServices represents all available docker compose services for the testing environment.
	TestServices = map[string]Service{
		"pg":          PostgresTestService,
		"postgres":    PostgresTestService,
		"postgresql":  PostgresTestService,
		"mysql":       MySQLTestService,
		"cockroach":   CockroachTestService,
		"cockroachdb": CockroachTestService,
		"redis":       RedisTestService,
	}
)
