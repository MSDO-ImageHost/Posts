package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func newPostHandler(consumer <-chan amqp.Delivery) {
	for msg := range consumer {
		if msg.Body == nil {
			log.Fatal("No data in message")
		}

		// Parse the received JSON into Post struct
		postReq := api.CreateRequest{}
		if err := json.Unmarshal(msg.Body, &postReq); err != nil {
			log.Println(err)
			return
		}

		newPost := storage.PostScaffold{
			AuthorID:      postReq.AuthToken,
			HeaderContent: postReq.Header,
			BodyContent:   postReq.Body,
		}

		log.Println(newPost)

		// Save in database
		_, err := storage.Posts.Add(newPost)

		// Acknowledge message processed
		removeFromQueue := true
		if err != nil {
			log.Println(err)
			removeFromQueue = false
		}
		if err := msg.Ack(removeFromQueue); err != nil {
			log.Fatal(err)
		}
	}
}
