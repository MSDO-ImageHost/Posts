package dao

import (
	"time"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoStorage stores the credentials to the database and collection
type MongoStorage struct {
	Host               *mongo.Database
	ScaffoldCollection *mongo.Collection
	HeaderCollection   *mongo.Collection
	BodyCollection     *mongo.Collection
}

type ScaffoldStorage MongoStorage
type HeaderStorage MongoStorage
type BodyStorage MongoStorage

// Database models
type PostQueryID string

// Content store field specific data for a post. It also contain references to revisions of itself
type Content struct {
	ID        primitive.ObjectID `bson:"_id"`
	AuthorID  string             `bson:"author_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Data      string             `bson:"data"`
	Replacing primitive.ObjectID `bson:"replacing,omitempty"`
	Meta      models.Meta        `bson:"-"`
}

// Post is the main frame for storing references to content in the post
type Scaffold struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	AuthorID  string             `bson:"author_id"`
	Header    primitive.ObjectID `bson:"header"`
	Body      primitive.ObjectID `bson:"body"`
	Meta      models.Meta        `bson:"-"`
}
