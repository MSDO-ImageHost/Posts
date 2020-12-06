package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
	jwt "github.com/dgrijalva/jwt-go"
)

var sig chan os.Signal = make(chan os.Signal)

func main() {

	log.Println(_LOG_TAG, "Starting up")

	var wg sync.WaitGroup
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
	if err := broker.NewSubPub(createSinglePostConf, createOnePostHandler); err != nil {
		log.Fatal(err)
	}

	// Read events
	if err := broker.NewSubPub(readSinglePostConf, readSinglePostHandler); err != nil {
		log.Fatal(err)
	}
	if err := broker.NewSubPub(readPostHistoryConf, readPostHistoryHandler); err != nil {
		log.Fatal(err)
	}
	if err := broker.NewSubPub(readUserPostsConf, readUserPostsHandler); err != nil {
		log.Fatal(err)
	}
	if err := broker.NewSubPub(readManyPostsConf, readManyPostsHandler); err != nil {
		log.Fatal(err)
	}

	// Update events
	if err := broker.NewSubPub(updateOnePostConf, updateOnePostHandler); err != nil {
		log.Fatal(err)
	}

	// Delete events
	if err := broker.NewSubPub(deleteOnePostConf, deleteOnePostHandler); err != nil {
		log.Fatal(err)
	}
	if err := broker.NewSubPub(deleteManyPostsConf, deleteManyPostsHandler); err != nil {
		log.Fatal(err)
	}

	log.Println(_LOG_TAG, "Up and running")
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func getToken(token *jwt.Token) (interface{}, error) {
	return os.Getenv("JWT_HMAC_SECRET"), nil
}
