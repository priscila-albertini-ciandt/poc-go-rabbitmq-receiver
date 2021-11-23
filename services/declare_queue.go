package services

import (
	"poc-rabbitmq/error"

	"github.com/streadway/amqp"
)

func DeclareQueue(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	error.FailOnError(err, "Failed to declare a queue")

	return q
}
