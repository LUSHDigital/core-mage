package compose

var (
	// PostgresService represents a docker compose postgresql service.
	PostgresService = Service{
		Image: "postgres",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"5432:5432",
		},
		Volumes: []string{
			"${PWD}/data/postgres/dev:/var/lib/postgresql/data",
		},
		Environment: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "passwd",
		},
		ExternalURLPattern: "%s:5432/service",
		InternalURLPattern: "%s:5432/service",
	}

	// PostgresTestService represents a docker compose postgresql service.
	PostgresTestService = Service{
		Image: "postgres",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"5433:5432",
		},
		Environment: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "passwd",
		},
		ExternalURLPattern: "%s:5433/service",
		InternalURLPattern: "%s:5432/service",
	}
)
