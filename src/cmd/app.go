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

	storedPostID, _ := storage.Posts.Add(newPost)
	fmt.Printf("INSERTING: %+v \t->\t%s\n", newPost, storedPostID)

	updatedPostID, _ := storage.Posts.Update(storedPostID, models.Post{
		Header: "Updated title",
	})
	fmt.Printf("UPDATING: %+v \t->\t%s\n", newPost, updatedPostID)

	fetchedPost, _ := storage.Posts.Find(storedPostID)
	fmt.Printf("FETCHING: %s \t -> \t %+v \n", storedPostID, fetchedPost)

	deletedPostID, _ := storage.Posts.Delete(storedPostID)
	fmt.Printf("REMOVING: %s\n", deletedPostID)

}
