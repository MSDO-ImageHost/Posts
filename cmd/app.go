package main

import (
	"fmt"
	"log"

	"github.com/MSDO-ImageHost/Posts/internal/database"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

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

	post := database.PostData{
		Author: "123-christian-id",
		Header: database.PostContent{Data: "Hello from header!"},
		Body:   database.PostContent{Data: "Hello from body!"},
	}
	newPost, err := storage.AddOnePost(post)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Inserted document")
	fmt.Println(PrettyFormatMap(newPost))

}
