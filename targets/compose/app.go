package compose

var (
	// AppService represents a docker compose app service.
	AppService = Service{
		Image:   "lushdigital/alpine-golang:latest",
		Command: "go run -mod=vendor service/main.go ",
		EnvFile: "infra/dev.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
			"$GOMODPATH:/go/pkg/mod",
		},
		WorkingDir: "/service",
	}
	// AppTestService represents a docker compose app service.
	AppTestService = Service{
		Image:   "lushdigital/alpine-golang:latest",
		Command: "go test -mod=vendor -v -cover ./...",
		EnvFile: "infra/test.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
			"$GOMODPATH:/go/pkg/mod",
		},
		WorkingDir: "/service",
	}
)
