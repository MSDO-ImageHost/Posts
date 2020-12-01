package main

import (
	"encoding/json"
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
	_ = database.PostData{
		Author: "123-christian-id",
		Header: database.PostContent{Data: "Hello from header!"},
		Body:   database.PostContent{Data: "Hello from body!"},
	}

	//storage.AddOnePost(post)

	res, err := storage.FindOnePost("5fc5937602ac8fa03917d22e")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(PrettyFormatMap(res))

}

func PrettyFormatMap(d interface{}) string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
