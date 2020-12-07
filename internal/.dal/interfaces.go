//go:generate mockgen -source interfaces.go -destination interfaces_mock.go -package database

package database

import "time"

//type MongoInterface interface {
//	NewPost(interface{}) (post StorageInterface, err error)
//}

type StorageInterface interface {
	// Getters
	GetID() string
	GetAuthorID() string
	GetCreatedAt() time.Time
	GetUpdatedAt() *time.Time
	GetMarkedDeleted() bool
	GetHeader() interface{}
	GetBody() interface{}

	Insert() error
	Find() error
	//UpdateHeader(update PostContent)
	//UpdateBody(update PostContent)
}

type ContentInterface interface {
	GetID() string
	GetAuthorID() string
	GetCreatedAt() time.Time
	GetData() string
	GetMarkedDeleted() bool
}
