package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func updateRequest(msg amqp.Delivery) (interface{}, bool, error) {

	// Parse the received JSON into Post struct
	postReq := api.UpdateRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return nil, false, err
	}

	updateData := storage.PostScaffold{
		IDHex:         postReq.PostID,
		HeaderContent: postReq.Header,
		BodyContent:   postReq.Body,
	}

	// Find and update in database
	result, err := storage.Posts.UpdateOne(updateData)
	if err != nil {
		log.Fatal(err)
	}

	postRes := api.UpdateResponse{
		PostID: result,
	}

	// Acknowledge message was processed
	if err != nil {
		return nil, false, err
	}
	return (interface{})(postRes), true, nil
}
