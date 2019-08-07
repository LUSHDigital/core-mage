package compose

var (
	// RedisService represents a docker compose redis service.
	RedisService = Service{
		Image:   "redis:4",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
		Volumes: []string{
			"${PWD}/data/redis/dev:/data",
		},
		Ports: []string{
			"6379:6379",
		},
	}

	// RedisTestService represents a docker compose redis service.
	RedisTestService = Service{
		Image:   "redis:4",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
	}
)
