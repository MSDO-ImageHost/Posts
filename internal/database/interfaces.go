//go:generate mockgen -source interfaces.go -destination interfaces_mock.go -package database

package database

type storageInterface interface {
	// Create
	AddOne(post PostData) (result PostData, err error)
	AddMany(posts []PostData) (results []PostData, err error)

	// Read
	FindOne(postIdHex string) (result PostData, err error)
	FindMany(postIdHexes []string) (results []PostData, err error)
	FindHistory(postIdHex string) (result PostDataHistory, err error)
	FindUserPosts(author string) (results []PostData, err error)

	// Update
	UpdateOne(post PostData) (result PostData, err error)

	// Delete
	DeleteOne(postIdHex string) (result PostData, err error)
	DeleteMany(postIdHexes []string) (results []PostData, err error)
}
