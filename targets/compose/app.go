package compose

var (
	// AppService represents a docker compose app service.
	AppService = Service{
		Build: Build{
			Context:    "${PWD}",
			Dockerfile: "${PWD}/Dockerfile",
		},
		Command: "go run -mod=vendor service/main.go ",
		EnvFile: "infra/dev.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
			"$GOMODPATH:/go/pkg/mod:ro",
		},
		WorkingDir: "/service",
	}
	// AppTestService represents a docker compose app service.
	AppTestService = Service{
		Build: Build{
			Context:    "${PWD}",
			Dockerfile: "${PWD}/Dockerfile",
		},
		Command: "go test -mod=vendor -v -cover ./...",
		EnvFile: "infra/test.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
			"$GOMODPATH:/go/pkg/mod:ro",
		},
		WorkingDir: "/service",
	}
)
