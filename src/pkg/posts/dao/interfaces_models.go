package dao

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/mongo"
)

type HeaderInterface ContentInterface
type BodyInterface ContentInterface

// MongoStorage stores the credentials to the database and collection
type MongoStorage struct {
	Host               *mongo.Database
	ScaffoldCollection *mongo.Collection
	HeaderCollection   *mongo.Collection
	BodyCollection     *mongo.Collection
}

type ScaffoldStorage MongoStorage
type HeaderStorage MongoStorage
type BodyStorage MongoStorage

// Public interface handlers for this package
var Posts ScaffoldInterface
var Headers HeaderInterface
var Bodies BodyInterface

type ScaffoldInterface interface {
	Add(post models.Post) (string, error)
	Update(postID string, post models.Post) (string, error)
	//Delete(postID string) (models.Post, error)
	Find(postID string) (models.Post, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}

type ContentInterface interface {
	Add(title models.Content) (string, error)
	//Delete(id models.PostQueryID) (models.Content, error)
	//Find(id models.PostQueryID) (models.Content, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}
