package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MSDO-ImageHost/Posts/pkg/api"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
	"github.com/streadway/amqp"
)

func getUserPostsRequest(msg amqp.Delivery) (interface{}, bool, error) {

	fmt.Println(msg.ConsumerTag)

	// Parse the received JSON into Post struct
	postReq := api.CreateRequest{}
	if err := json.Unmarshal(msg.Body, &postReq); err != nil {
		log.Println(err)
		return nil, false, err
	}

	userId := postReq.AuthToken

	// Find in database
	results, err := storage.Posts.FindUserPosts(userId)
	if err != nil {
		log.Fatal(err)
	}

	postRes := make([]api.GetResponse, len(results))
	for i, result := range results {
		postRes[i] = api.GetResponse{
			PostID:    result.IDHex,
			CreatedAt: result.CreatedAt,
			AuthorID:  result.AuthorID,
			Header:    result.HeaderContent,
			Body:      result.BodyContent,
		}
	}

	// Acknowledge message was processed
	if err != nil {
		return nil, false, err
	}
	return (interface{})(postRes), true, nil
}
