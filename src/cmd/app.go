package main

import (
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
		Meta: models.Meta{
			TZ: "Europe/Copenhagen",
		},
	}

	storage.Posts.Add(newPost)

}
