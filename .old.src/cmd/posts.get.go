package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func getRequest(msg amqp.Delivery) (interface{}, bool, error) {

	// Parse the received JSON into Post struct
	postReq := api.GetRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return nil, false, err
	}

	// Find in database
	result, err := storage.Posts.FindOne(postReq.PostID)
	if err != nil {
		log.Fatal(err)
	}

	postRes := api.GetResponse{
		PostID:    result.ID.Hex(),
		CreatedAt: result.CreatedAt,
		AuthorID:  result.AuthorID,
		Header:    result.HeaderContent,
		Body:      result.BodyContent,
	}

	// Acknowledge message was processed
	if err != nil {
		return nil, false, err
	}
	return (interface{})(postRes), true, nil
}
