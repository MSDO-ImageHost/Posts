package dao

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
)

// Update
func (db *MongoStorage) Update(models.PostQueryID) (models.Post, error) {
	return models.Post{}, nil
}
