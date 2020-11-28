package main

import (
	"log"
	"os"

	broker "github.com/MSDO-ImageHost/Posts/pkg/broker"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
)

var forever chan bool = make(chan bool)

func main() {

	log.Println("App: Starting posts app")

	// Initialize storage and deinitialize on exit
	log.Println("App: Initializing database")
	if err := storage.Init(); err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := storage.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("App: Database initialized")

	// Initialize message broker and deinitialize on exit
	log.Println("App: Initializing broker")
	if err := broker.Init(); err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := broker.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("App: Broker initialized")

	// Create event handlers
	broker.Broker.NewAsyncHandler("posts.create", createRequest)
	broker.Broker.NewAsyncHandler("posts.get", getRequest)
	broker.Broker.NewAsyncHandler("posts.update", updateRequest)
	broker.Broker.NewAsyncHandler("posts.delete", deleteRequest)
	broker.Broker.NewAsyncHandler("posts.get.many", getManyRequest)
	broker.Broker.NewAsyncHandler("posts.get.history", getHistoryRequest)
	broker.Broker.NewAsyncHandler("posts.get.userposts", getUserPostsRequest)

	// Run application 'forever'
	<-forever
	os.Exit(0)
}

func JWTTokenHandler(token string) (bool, error) {
	return true, nil
}
