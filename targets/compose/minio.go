package compose

var (
	// MinioService represents a docker compose minio service.
	MinioService = Service{
		Image:   "minio/minio",
		Command: "server /data",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
		Volumes: []string{
			"${PWD}/data/minio/dev:/data",
		},
		Ports: []string{
			"9000:9000",
		},
		Environment: map[string]string{
			"MINIO_ACCESS_KEY": "miniouser",
			"MINIO_SECRET_KEY": "miniopasswd",
		},
		ExternalURLPattern: "miniouser:miniopasswd@%s:9000",
		InternalURLPattern: "miniouser:miniopasswd@%s:9000",
	}
	// MinioTestService represents a docker compose minio service for the testing environment.
	MinioTestService = Service{
		Image:   "minio/minio",
		Command: "server /data",
		Restart: "always",
		Logging: map[string]string{
			"driver": "none",
		},
		Ports: []string{
			"9001:9000",
		},
		Environment: map[string]string{
			"MINIO_ACCESS_KEY": "miniouser",
			"MINIO_SECRET_KEY": "miniopasswd",
		},
		ExternalURLPattern: "miniouser:miniopasswd@%s:9001",
		InternalURLPattern: "miniouser:miniopasswd@%s:9000",
	}
)
