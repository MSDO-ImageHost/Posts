package database

type Interface interface {

	// Create methods
	Add(post PostScaffold) (postId string, err error)

	// Read methods
	FindOne(postIDHex string) (foundPost PostScaffold, err error)
	FindOneHistory(postIDHex string) (postHistory PostScaffold, err error)
	FindUserPosts(userId string) (userPosts []PostScaffold, err error)

	// Update methods
	UpdateOne(post PostScaffold) (updatedPostID string, err error)
	//UpdateMany() ([]string, err error)

	// Delete methods
	DeleteOne(postIDHex string) (postId string, err error)
	//DeleteMany(postIDs []string) (string[], err error)
	//FindMany() ([]PostScaffold, err error)
}
