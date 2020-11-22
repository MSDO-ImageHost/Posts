package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/MSDO-ImageHost/Posts/pkg/broker"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
)

func newPostHandler(queueName string) {
	queue, err := broker.GetBroker().Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := broker.GetBroker().Channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	func() {
		for msg := range consumer {

			if msg.Body == nil {
				log.Fatal("No data in message")
			}

			// Parse the received JSON into Post struct
			postReq := api.CreateRequest{}
			if json.Unmarshal(msg.Body, &postReq); err != nil {
				log.Println(err)
				return
			}

			newPost := storage.PostScaffold{
				AuthorID:      postReq.AuthToken,
				HeaderContent: postReq.Header,
				BodyContent:   postReq.Body,
			}

			_, err := storage.Posts.Add(newPost)

			if err != nil {
				log.Println(err)
				msg.Ack(false)
			}
			msg.Ack(true)
		}
	}()
}
