package main

import (
	"log"

	broker "github.com/MSDO-ImageHost/Posts/pkg/broker"
	storage "github.com/MSDO-ImageHost/Posts/pkg/database"
)

func main() {

	// Initialize storage and deinitialize on exit
	if err := storage.Init(); err != nil {
		log.Panicln(err)
	}
	defer func() {
		if err := storage.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize message broker
	if err := broker.Init(); err != nil {
		log.Panicln(err)
	}
	// Deinitialize on exit
	defer func() {
		if err := broker.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()

	go newPostHandler("posts.create")
	forever := make(chan bool)
	<-forever
}
