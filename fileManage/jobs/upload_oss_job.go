package jobs

import (
	"encoding/json"
	"fileManage/config"
	"fileManage/databases"
	"fileManage/rabbitMQ/RabbitMQ"
	"fileManage/services"
	"fileManage/util"
	"fmt"
	"log"
	"path/filepath"
)

const queueName = "uploadOssJob"

type UploadOssJobMessage struct {
	MediaId uint `json:"media_id"`
}

func NewPublishUploadOssJob(uploadOssJobMessage UploadOssJobMessage) {
	rabbitmq := RabbitMQ.NewRabbitMQSimple(queueName)
	uploadOssJobMessageStr, _ := json.Marshal(uploadOssJobMessage)
	rabbitmq.PublishSimple(string(uploadOssJobMessageStr))
	fmt.Println("发送成功！")
}

func NewConsumeSimpleUploadOssJob() {
	r := RabbitMQ.NewRabbitMQSimple(queueName)
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.Channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//接收消息
	msgs, err := r.Channel.Consume(
		q.Name, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			//消息逻辑处理，可以自行设计逻辑
			log.Printf("Received a message: %s", d.Body)
			uploadOssJobMessage := &UploadOssJobMessage{}
			err := json.Unmarshal(d.Body, uploadOssJobMessage)
			if err != nil {
				return
			}
			if uploadOssJobMessage != nil && uploadOssJobMessage.MediaId > 0 {
				media, err2 := databases.GetMediaById(uploadOssJobMessage.MediaId)
				if err2 != nil {
					return
				}
				if media.OssPath != "" {
					return
				}
				ossFileName, err := services.OssUpload(config.AppConf.StorageConf.Path+util.GetPathTag()+media.Path, media.Md5+filepath.Ext(media.Name))
				if err != nil {
					return
				}
				media.OssPath = ossFileName
				err = databases.UpdateMediaByMedia(media)
				if err != nil {
					return
				}
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
