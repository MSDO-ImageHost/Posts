package main

import (
	"encoding/json"
	"log"

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

}

func PrettyFormatMap(d interface{}) string {
	b, err := json.MarshalIndent(d, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
