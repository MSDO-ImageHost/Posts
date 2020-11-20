package dao

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
)

// Public interface handlers for this package
var Posts ScaffoldInterface
var Headers HeaderInterface
var Bodies BodyInterface

type ScaffoldInterface interface {
	Add(post models.Post) (string, error)
	//Update(id models.PostQueryID) (string, error)
	//Delete(id models.PostQueryID) (models.Post, error)
	//Find(id models.PostQueryID) (models.Post, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}

type ContentInterface interface {
	Add(title Content) (string, error)
	//Update(id string) (models.Post, error)
	//Delete(id models.PostQueryID) (models.Content, error)
	//Find(id models.PostQueryID) (models.Content, error)
	//FindMany(ids []models.PostQueryID) ([]models.Post, error)
}

type HeaderInterface ContentInterface
type BodyInterface ContentInterface
