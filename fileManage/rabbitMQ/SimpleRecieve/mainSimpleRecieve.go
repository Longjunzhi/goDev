package main

import (
	"Img/rabbitMQ/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.ConsumeSimple()
}
