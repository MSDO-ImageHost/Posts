package dao

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
)

// Find
func (db *MongoStorage) Find(models.PostQueryID) (models.Post, error) {
	return models.Post{}, nil
}
