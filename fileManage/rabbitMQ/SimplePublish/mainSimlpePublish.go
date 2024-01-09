package main

import (
	"fileManage/rabbitMQ/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
		"kuteng")
	rabbitmq.PublishSimple("Hello kuteng111!")
	fmt.Println("发送成功！")
}
