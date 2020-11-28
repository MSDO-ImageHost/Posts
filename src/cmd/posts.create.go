package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func createRequest(msg amqp.Delivery) (interface{}, bool, error) {

	// Parse the received JSON into Post struct
	postReq := api.CreateRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		return nil, false, err
	}

	newPost := storage.PostScaffold{
		AuthorID:      postReq.AuthToken,
		HeaderContent: postReq.Header,
		BodyContent:   postReq.Body,
	}

	// Save in database
	result, err := storage.Posts.Add(newPost)
	if err != nil {
		log.Fatal(err)
	}

	postRes := api.CreateResponse{
		PostID: result,
	}

	// Acknowledge message was processed
	if err != nil {
		return nil, false, err
	}
	return (interface{})(postRes), true, nil
}
