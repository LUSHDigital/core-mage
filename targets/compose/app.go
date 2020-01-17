package compose

var (
	// AppService represents a docker compose app service.
	AppService = Service{
		Command: "go run -mod=vendor service/main.go",
		EnvFile: "infra/compose.dev.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
		},
		WorkingDir: "/service",
	}
	// AppTestService represents a docker compose app service.
	AppTestService = Service{
		Command: "go test -mod=vendor -v -cover ./...",
		EnvFile: "infra/compose.test.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
		},
		WorkingDir: "/service",
	}
)
