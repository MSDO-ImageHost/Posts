package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func updatePostHandler(msg amqp.Delivery) bool {

	// Parse the received JSON into Post struct
	postReq := api.UpdateRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return false
	}

	// Find and update in database
	//_, err := storage.Posts.UpdateOne(postReq.PostID)

	// Acknowledge message was processed
	//if err != nil {
	//	return false
	//}
	return true

}
