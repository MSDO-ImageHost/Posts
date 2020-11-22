package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func updatePostHandler(consumer <-chan amqp.Delivery) {

	for msg := range consumer {
		if msg.Body == nil {
			log.Fatal("No data in message")
		}

		// Parse the received JSON into Post struct
		postReq := api.UpdateRequest{}
		if err := json.Unmarshal(msg.Body, &postReq); err != nil {
			log.Println(err)
			return
		}

		// Find and update in database
		/*_, err := storage.Posts.UpdateOne(postReq.PostID)

		// Acknowledge message processed
		removeFromQueue := true
		if err != nil {
			log.Println(err)
			removeFromQueue = false
		}
		if err := msg.Ack(removeFromQueue); err != nil {
			log.Fatal(err)
		}
		*/
	}
}
