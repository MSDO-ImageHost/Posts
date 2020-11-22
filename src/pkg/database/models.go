package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoShell struct {
	Client *mongo.Client
	DB     *mongo.Database
}

type MongoStorage struct {
	ScaffoldStorage *mongo.Collection
	HeaderStorage   *mongo.Collection
	BodyStorage     *mongo.Collection
}

type PostQueryID string

// Scaffold is the main frame for storing references to content in the post
type PostScaffold struct {
	ID            primitive.ObjectID   `bson:"_id"`
	CreatedAt     time.Time            `bson:"created_at"`
	AuthorID      string               `bson:"author_id"`
	HeaderContent string               `bson:"-"`
	Header        []primitive.ObjectID `bson:"header_ids"`
	BodyContent   string               `bson:"-"`
	Body          []primitive.ObjectID `bson:"body_ids"`
}

type Content struct {
	ID        primitive.ObjectID `bson:"_id"`
	AuthorID  string             `bson:"author_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Data      string             `bson:"data"`
}
