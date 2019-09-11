package compose

var (
	// AppService represents a docker compose app service.
	AppService = Service{
		Command: "go run -mod=vendor service/main.go",
		EnvFile: "infra/dev.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
		},
		WorkingDir: "/service",
	}
	// AppTestService represents a docker compose app service.
	AppTestService = Service{
		Command: "go test -mod=vendor -v -cover ./...",
		EnvFile: "infra/test.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
		},
		WorkingDir: "/service",
	}
)
