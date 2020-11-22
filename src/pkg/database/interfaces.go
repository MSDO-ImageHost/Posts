package database

type Interface interface {
	Add(post PostScaffold) (string, error)
	//FindOne(postID string) (PostScaffold, error)
	//UpdateOne() (string, error)
	//UpdateMany() ([]string, error)
	//DeleteOne() (string[], error)
	//DeleteMany() (string[], error)
	//FindMany() ([]storage.PostScaffold, error)
	//FindOneHistory() (, error)
}
