//go:generate mockgen -source interfaces.go -destination interfaces_mock.go -package database

package database

import (
	"github.com/MSDO-ImageHost/Posts/internal/api"
	"github.com/MSDO-ImageHost/Posts/internal/auth"
)

type storageInterface interface {
	// Create
	AddOne(post PostData) (result PostData, err error)

	// Read
	FindOne(postIdHex string) (result PostData, err error)
	FindMany(postIdHexes []string, paging api.PagingStruct) (results []PostData, err error)
	FindHistory(postIdHex string) (result PostHistoryData, err error)
	FindUserPosts(author string, paging api.PagingStruct) (results []PostData, err error)

	// Update
	UpdateOne(post PostData, auth auth.User) (result PostData, err error)

	// Delete
	DeleteOne(postIdHex string, auth auth.User) (result PostData, err error)
	DeleteMany(postIdHexes []string, auth auth.User) (results []string, err error)
	MarkDeleteOne(postIdHex string, auth auth.User) (result string, err error)
}
