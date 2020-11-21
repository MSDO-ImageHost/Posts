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
	Add(post models.Post) (createdPostID string, err error)
	Update(postID string, post models.Post) (updatedPostID string, err error)
	Delete(postID string) (deletedPostID string, err error)
	Find(postID string) (fetchedPost models.Post, err error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}

type ContentInterface interface {
	Add(content models.Content) (contentID string, err error)
	//Delete(id models.PostQueryID) (models.Content, error)
	//Find(id models.PostQueryID) (models.Content, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}
