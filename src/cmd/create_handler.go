package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func newPostHandler(msg amqp.Delivery) bool {

	// Parse the received JSON into Post struct
	postReq := api.CreateRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return false
	}

	newPost := storage.PostScaffold{
		AuthorID:      postReq.AuthToken,
		HeaderContent: postReq.Header,
		BodyContent:   postReq.Body,
	}

	log.Println(newPost)

	// Save in database
	_, err := storage.Posts.Add(newPost)

	// Acknowledge message was processed
	if err != nil {
		return false
	}
	return true
}
