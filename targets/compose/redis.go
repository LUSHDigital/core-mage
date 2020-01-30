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
		ExternalURLPattern: "%s:6379/0",
		InternalURLPattern: "%s:6379/0",
	}

	// RedisTestService represents a docker compose redis service.
	RedisTestService = Service{
		Image:   "redis:4",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
		Ports: []string{
			"6380:6379",
		},
		ExternalURLPattern: "%s:6380/0",
		InternalURLPattern: "%s:6379/0",
	}
)
