package dao

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
)

// Delete
func (db *MongoStorage) Delete(PostQueryID) (models.Post, error) {
	return models.Post{}, nil
}
