package receive

import (
	"poc-rabbitmq/config"
	"poc-rabbitmq/services"
)

func Receive() {

	conn := config.Connect()
	defer conn.Close()

	ch := services.OpenChannel(conn)
	defer ch.Close()

	q := services.DeclareQueue(ch, "queue.hello")

	services.Consume(ch, q)
}
