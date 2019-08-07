package compose

var (
	// AppTestService represents a docker compose app service.
	AppTestService = Service{
		Image:   "lushdigital/alpine-golang:latest",
		Command: "go test -mod=vendor -v -cover ./...",
		EnvFile: "infra/test.env",
		Restart: "no",
		Volumes: []string{
			"$PWD:/service:ro",
		},
		WorkingDir: "/service",
	}
)
