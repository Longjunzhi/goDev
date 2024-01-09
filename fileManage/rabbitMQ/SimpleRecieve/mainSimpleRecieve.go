package main

import (
	"fileManage/rabbitMQ/RabbitMQ"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.ConsumeSimple()
}
