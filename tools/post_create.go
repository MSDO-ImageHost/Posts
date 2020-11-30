package main

import (
	"fmt"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/broker"
	"github.com/streadway/amqp"
)

var forever chan bool = make(chan bool)

func main() {

	conn, err := amqp.Dial("amqp://ImageHostPosts:DM8742020@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Declare queue
	queue, err := ch.QueueDeclare(
		"gateway.replies",
		broker.QConfig.Durable,
		broker.QConfig.AutoDelete,
		broker.QConfig.Exclusive,
		broker.QConfig.NoWait,
		broker.QConfig.Args,
	)
	if err != nil {
		log.Fatal(err)
	}

	/** Listen for responses **/
	consume, err := ch.Consume(
		queue.Name,
		broker.CConfig.Consumer,
		broker.CConfig.AutoAck,
		broker.CConfig.Exclusive,
		broker.CConfig.NoLocal,
		broker.CConfig.NoWait,
		broker.CConfig.Args,
	)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for msg := range consume {
			fmt.Println(string(msg.Body))
		}
	}()
	/** **/

	/** Publish payload **/
	postPayload := amqp.Publishing{
		ContentType: "application/json",
		Body: []byte(`{
			"auth_token": "<JWT>",
			"header": "<String: title of the post>",
			"body": "<String: body text of the post>"
		}`),
		ReplyTo: "gateway.replies",
	}

	if err := ch.Publish("", "posts-create", false, false, postPayload); err != nil {
		log.Fatal(err)
	}
	/** **/

	<-forever
}
