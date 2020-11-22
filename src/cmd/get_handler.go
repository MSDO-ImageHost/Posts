package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func getPostHandler(msg amqp.Delivery) bool {

	// Parse the received JSON into Post struct
	postReq := api.GetRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return false
	}

	// Find in database
	_, err := storage.Posts.FindOne(postReq.PostID)

	// Acknowledge message was processed
	if err != nil {
		return false
	}
	return true
}
