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
			"26257:26257",
			"5432:5432",
		},
		Volumes: []string{
			"${PWD}/data/postgres/dev:/var/lib/postgresql/data",
		},
		Environment: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "passwd",
		},
	}

	// PostgresTestService represents a docker compose postgresql service.
	PostgresTestService = Service{
		Image: "postgres",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Volumes: []string{
			"${PWD}/data/postgres/test:/var/lib/postgresql/data",
		},
		Environment: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "passwd",
		},
	}
)
