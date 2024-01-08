package main

import (
	"Img/rabbitMQ/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.PublishSimple("Hello kuteng222!")
	fmt.Println("发送成功！")
}
