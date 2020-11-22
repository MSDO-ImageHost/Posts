package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func getManyPostsHandler(msg amqp.Delivery) bool {

	// Parse the received JSON into Post struct
	postReq := api.GetManyRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return false
	}

	// Find in database
	//_, err := storage.Posts.FindMany(postReq.PostID)

	// Acknowledge message was processed
	//if _ != nil {
	//	return false
	//}
	return true

}
