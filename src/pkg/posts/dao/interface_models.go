package dao

import models "github.com/MSDO-ImageHost/Posts/pkg/posts"

// MongoDBHost stores the credentials to the database and collection
type MongoDBHost struct {
	URI  string
	Name string
}

type MongoCollection struct {
	Host       MongoDBHost
	Collection string
}

type ScaffoldStorage MongoCollection
type HeaderStorage MongoCollection
type BodyStorage MongoCollection

// Public interface handlers for this package
var Posts PostStorage
var Titles ContentStorage
var Bodies ContentStorage

// PostStorage interface options
type PostStorage interface {
	Add(post models.Post) (models.Post, error)
	Update(id models.PostQueryID) (models.Post, error)
	Delete(id models.PostQueryID) (models.Post, error)
	Find(id models.PostQueryID) (models.Post, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}

type ContentStorage interface {
	//Add(title models.Content) (models.Content, error)
	//Update(id string) (models.Post, error)
	//Delete(id models.PostQueryID) (models.Content, error)
	//Find(id models.PostQueryID) (models.Content, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}
