package main

import (
	"fmt"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	storage "github.com/MSDO-ImageHost/Posts/pkg/posts/dao"
)

func main() {

	err := storage.Init()
	if err != nil {
		panic(err)
	}

	newPost := models.Post{
		AuthorID: "123-christian-id",
		Header:   "Hello title!",
		Body:     "This is my first post..!",
	}

	fmt.Printf("Created post:\n%+v\n", newPost)
	storedPostID, _ := storage.Posts.Add(newPost)
	fmt.Printf("Inserted it and received id %s\n\n", storedPostID)

	fmt.Printf("Fetching it again %s\n", storedPostID)
	fetchedPost, err := storage.Posts.Find(storedPostID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Fetched:\n%+v\n", fetchedPost)
}
