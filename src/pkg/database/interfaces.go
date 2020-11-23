package database

type Interface interface {
	Add(post PostScaffold) (string, error)
	FindOne(postID string) (PostScaffold, error)
	Update(postID string, post PostScaffold) (string, error)
	//UpdateMany() ([]string, error)
	DeleteOne(postID string) (string, error)
	//DeleteMany(postIDs []string) (string[], error)
	//FindMany() ([]storage.PostScaffold, error)
	//FindOneHistory() (, error)
}
