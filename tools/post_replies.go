package main

import (
	"fmt"
	"log"

	"github.com/MSDO-ImageHost/Posts/internal/utils"
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
		"confirm-post-creation",
		true,
		true,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	/** Listen for responses **/
	consume, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range consume {
			fmt.Println(utils.PrettyFormatMap(msg.Body))
		}
	}()

	<-forever
}
