package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func deletePostHandler(consumer <-chan amqp.Delivery) {

	for msg := range consumer {
		if msg.Body == nil {
			log.Fatal("No data in message")
		}

		// Parse the received JSON into Post struct
		postReq := api.DeleteRequest{}
		if err := json.Unmarshal(msg.Body, &postReq); err != nil {
			log.Println(err)
			return
		}

		// Find in database and delete
		/*_, err := storage.Posts.DeleteOne(postReq.PostID)

		// Acknowledge message processed
		removeFromQueue := true
		if err != nil {
			log.Println(err)
			removeFromQueue = false
		}
		if err := msg.Ack(removeFromQueue); err != nil {
			log.Fatal(err)
		}*/
	}
}
