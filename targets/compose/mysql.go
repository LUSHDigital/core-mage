package compose

var (
	// MySQLService represents a docker compose mysql service.
	MySQLService = Service{
		Image: "mysql:5.7",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"3306:3306",
		},
		Volumes: []string{
			"${PWD}/data/mysql/dev:/var/lib/mysql",
		},
		Environment: map[string]string{
			"MYSQL_DATABASE":      "service",
			"MYSQL_USER":          "user",
			"MYSQL_PASSWORD":      "passwd",
			"MYSQL_ROOT_PASSWORD": "passwd",
		},
		ExternalURLPattern: "tcp(%s:3306)/service",
		InternalURLPattern: "tcp(%s:3306)/service",
	}

	// MySQLTestService represents a docker compose mysql service for the test environment.
	MySQLTestService = Service{
		Image: "mysql:5.7",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"3307:3306",
		},
		Environment: map[string]string{
			"MYSQL_DATABASE":      "service",
			"MYSQL_USER":          "user",
			"MYSQL_PASSWORD":      "passwd",
			"MYSQL_ROOT_PASSWORD": "passwd",
		},
		ExternalURLPattern: "tcp(%s:3307)/service",
		InternalURLPattern: "tcp(%s:3306)/service",
	}
)
