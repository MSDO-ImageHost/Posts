package dao

import (
	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
)

// Delete
func (db *ScaffoldStorage) Delete(postID string) (models.Post, error) {
	return models.Post{}, nil
}
