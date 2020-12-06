package main

import (
	"fmt"
	"log"
	"os"

	dal "github.com/MSDO-ImageHost/Posts/internal/dal"
	"github.com/MSDO-ImageHost/Posts/internal/utils"
)

func main() {

	if err := dal.Connect(os.Getenv("MONGO_CONN_URI")); err != nil {
		log.Fatal(err)
	}

	data := make(map[string]interface{})
	data["creator_id"] = "123-christian-id"
	data["header"] = "Hello from header!"
	data["body"] = "This is the body"

	post, err := dal.NewPost(data)
	if err != nil {
		log.Fatal(err)
	}
	if err := post.Insert(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(utils.PrettyFormatMap(post))
	fmt.Print("\n\n\n")

	data2 := make(map[string]interface{})
	data2["_id"] = post.GetID()
	searchPost, err := dal.NewPost(data2)
	if err != nil {
		log.Fatal(err)
	}
	if err := searchPost.Find(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(utils.PrettyFormatMap(searchPost))

	if err := dal.Disconnect(); err != nil {
		log.Fatal(err)
	}

}
