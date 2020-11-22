package main

import (
	"log"

	broker "github.com/MSDO-ImageHost/Posts/pkg/broker"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

var forever chan bool = make(chan bool)

func main() {

	// Initialize storage and deinitialize on exit
	if err := storage.Init(); err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := storage.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize message broker and deinitialize on exit
	if err := broker.Init(); err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := broker.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create event handlers
	newEventHandler("posts.create", newPostHandler)
	newEventHandler("posts.get", getPostHandler)
	newEventHandler("posts.update", updatePostHandler)
	newEventHandler("posts.delete", deletePostHandler)
	newEventHandler("posts.get.many", getManyPostsHandler)
	newEventHandler("posts.get.history", getPostHistoryHandler)

	// Run application 'forever'
	<-forever
}

func newEventHandler(queueName string, handler func(amqp.Delivery) bool) {
	queue, err := broker.GetBroker().Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := broker.GetBroker().Channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for msg := range consumer {
			if msg.Body == nil {
				log.Println("No data in message")
			}
			// Acknowledge message processed
			msg.Ack(handler(msg))
		}
	}()
}
