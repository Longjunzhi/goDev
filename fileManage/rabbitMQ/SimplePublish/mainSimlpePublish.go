package main

import (
	"Img/rabbitMQ/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.PublishSimple("Hello kuteng111!")
	fmt.Println("发送成功！")
}
