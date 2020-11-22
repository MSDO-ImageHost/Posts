package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func deletePostHandler(msg amqp.Delivery) bool {

	// Parse the received JSON into Post struct
	postReq := api.DeleteRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return false
	}

	// Find in database and delete
	//_, err := storage.Posts.DeleteOne(postReq.PostID)

	// Acknowledge message processed
	//if err != nil {
	//	return false
	//}
	return true
}
