package compose

var (
	// RabbitMQService represents a docker compose RabbitMQ service.
	RabbitMQService = Service{
		Image: "rabbitmq:3.5.7-management",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"15672:15672",
			"5672:5672",
		},
		ExternalURLPattern: "amqp://guest@%s:5672",
		InternalURLPattern: "amqp://guest@%s:5672",
	}

	// RabbitMQTestService represents a test-time docker compose RabbitMQ service.
	RabbitMQTestService = Service{
		Image: "rabbitmq:3.5.7-management",
		Logging: map[string]string{
			"driver": "none",
		},
		Restart: "always",
		Ports: []string{
			"15673:15672",
			"5673:5672",
		},
		ExternalURLPattern: "amqp://guest@%s:5673",
		InternalURLPattern: "amqp://guest@%s:5673",
	}
)
