package main

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	storage "github.com/MSDO-ImageHost/Posts/pkg/posts/dao"
)

func main() {

	err := storage.Setup()
	if err != nil {
		panic(err)
	}

	newPost := models.Post{
		Author: "123-christian-id",
		Title: models.Content{
			Author: "123-christian-id",
			Data:   "Hello title!",
		},
		Body: models.Content{
			Author: "123-christian-id",
			Data:   "This is my first post..!",
		},
	}

	storage.Posts.Add(newPost)

}
