package compose

var (
	// FirebaseService represents a docker compose firebase service.
	FirebaseService = Service{
		Image:   "wceolin/firebase-emulator",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
		Ports: []string{
			"9898:9000",
		},
		ExternalURLPattern: "//%s:9898/0",
		InternalURLPattern: "//%s:9898/0",
		Command:            "firebase emulators:start --only database",
	}

	// FirebaseTestService represents a docker compose firebase service.
	FirebaseTestService = Service{
		Image:   "wceolin/firebase-emulator",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
		Ports: []string{
			"9899:9000",
		},
		ExternalURLPattern: "//%s:9899/0",
		InternalURLPattern: "//%s:9899/0",
		Command:            "firebase emulators:start --only database",
	}
)
