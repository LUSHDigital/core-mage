package compose

var (
	// RabbitMQService represents a docker compose RabbitMQ service.
	RabbitMQService = Service{
		Image: "rabbitmq:3-management",
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
)
