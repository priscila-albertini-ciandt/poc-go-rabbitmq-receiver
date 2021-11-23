package services

import (
	"poc-rabbitmq/error"

	"github.com/streadway/amqp"
)

func OpenChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	error.FailOnError(err, "Failed to open a channel")
	return ch
}
