package compose

var (
	// CockroachService represents a docker compose cockroach db service.
	CockroachService = Service{
		Image:   "cockroachdb/cockroach:v19.1.0",
		Command: "start --insecure --listen-addr 0.0.0.0:26257",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"26257:26257",
			"8080:8080",
		},
		Volumes: []string{
			"${PWD}/data/cockroach/dev1:/cockroach/cockroach-data",
		},
		urlPattern: "root@%s:26257/defaultdb?sslmode=disable",
	}

	// CockroachTestService represents a docker compose cockroach db service.
	CockroachTestService = Service{
		Image:   "cockroachdb/cockroach:v19.1.0",
		Command: "start --insecure --listen-addr 0.0.0.0:26257",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart:    "always",
		urlPattern: "root@%s:26257/defaultdb?sslmode=disable",
	}
)
