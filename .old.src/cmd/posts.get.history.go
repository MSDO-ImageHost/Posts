package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func getHistoryRequest(msg amqp.Delivery) (interface{}, bool, error) {

	// Parse the received JSON into Post struct
	postReq := api.GetHistoryRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return nil, false, err
	}

	// Find in database
	result, err := storage.Posts.FindOneHistory(postReq.PostID)
	if err != nil {
		log.Fatal(err)
	}

	postRes := api.UpdateResponse{
		PostID: result.IDHex,
	}

	// Acknowledge message was processed
	if err != nil {
		return nil, false, err
	}
	return (interface{})(postRes), true, nil

}
