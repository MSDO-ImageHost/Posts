package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
	"github.com/MSDO-ImageHost/Posts/internal/utils"
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
	log.Println(_LOG_TAG, "Registrering post creation events handler")
	if err := broker.NewSubPub(createSinglePostConf, createOnePostHandler); err != nil {
		log.Fatal(err)
	}

	// Read events
	log.Println(_LOG_TAG, "Registrering post read events handler")
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
	log.Println(_LOG_TAG, "Registrering post update events handler")
	if err := broker.NewSubPub(updateOnePostConf, updateOnePostHandler); err != nil {
		log.Fatal(err)
	}

	// Delete events
	log.Println(_LOG_TAG, "Registrering post delete events handler")
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

func handlerPlaceholder(req broker.HandleRequestPayload) (res broker.HandleResponsePayload, err error) {
	log.Println(_LOG_TAG, "placeholder", utils.PrettyFormatMap(req))
	return res, nil
}
