package main

import (
	"log"
	"sync"

	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

var forever chan bool = make(chan bool)

var wg sync.WaitGroup

func main() {

	log.Println(_LOG_TAG, "Starting up")
	wg.Add(2)

	// Initialize storage and deinitialize on exit
	go func() {
		log.Println(_LOG_TAG, "Initializing database")
		if err := storage.Init(); err != nil {
			log.Panicln(err)
		}
		wg.Done()
	}()
	defer func() {
		if err := storage.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize storage and deinitialize on exit
	go func() {
		log.Println(_LOG_TAG, "Initializing broker")
		if err := broker.Init(); err != nil {
			log.Panicln(err)
		}
		wg.Done()
	}()
	defer func() {
		if err := broker.Deinit(); err != nil {
			log.Fatal(err)
		}
	}()
	wg.Wait()

	// Create events
	log.Println(_LOG_TAG, "Registrering post creation events handler")
	broker.NewSubPub(postCreationHandleConf, postCreationHandler)

	// Read events
	log.Println(_LOG_TAG, "Registrering post read events handler")
	//broker.NewSubPub(postCreationHandleConf, postCreationHandler)

	// Update events
	log.Println(_LOG_TAG, "Registrering post update events handler")
	//broker.HandleUpdateOnePostEvents()

	// Delete events
	log.Println(_LOG_TAG, "Registrering post delete events handler")
	//broker.NewSubPub(postCreationHandleConf, postCreationHandler)

	log.Println(_LOG_TAG, "Up and running")
	<-forever
}
