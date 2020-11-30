package main

import (
	"encoding/json"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	"github.com/streadway/amqp"
)

func getManyRequest(msg amqp.Delivery) (interface{}, bool, error) {

	// Parse the received JSON into Post struct
	postReq := api.GetManyRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return nil, false, err
	}

	// Find in database
	//	_, err := storage.Posts.FindMany(postReq.PostID)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	postRes := api.UpdateResponse{
	//		PostID: result.ID.Hex(),
	//	}
	//
	//	// Acknowledge message was processed
	//	if err != nil {
	//		return nil, false, err
	//	}
	//	return (interface{})(postRes), true, nil
	return (interface{})(nil), true, nil

}
