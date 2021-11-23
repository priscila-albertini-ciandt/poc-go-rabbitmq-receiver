package config

import (
	"poc-rabbitmq/error"

	"github.com/streadway/amqp"
)

func Connect() *amqp.Connection {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	error.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}
