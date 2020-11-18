package dao

import (
	"context"
	"fmt"
	"time"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Add
func (db *ScaffoldStorage) Add(post models.Post) (models.Post, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.Host.URI))
	if err != nil {
		return post, nil
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	collection := client.Database(db.Host.Name).Collection(db.Collection)
	res, err := collection.InsertOne(ctx, post)
	if err != nil {
		return post, nil
	}
	fmt.Println(res.InsertedID)

	return post, nil
}

// Update
func (db *ScaffoldStorage) Update(models.PostQueryID) (models.Post, error) {
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

	return newPost, nil
}

// Delete
func (db *ScaffoldStorage) Delete(models.PostQueryID) (models.Post, error) {
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

	return newPost, nil
}

// Find
func (db *ScaffoldStorage) Find(models.PostQueryID) (models.Post, error) {
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

	return newPost, nil
}
