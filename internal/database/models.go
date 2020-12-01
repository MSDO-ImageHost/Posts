package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoShell struct {
	Client *mongo.Client
	DB     *mongo.Database
}

type mongoStorage struct {
	ScaffoldStorage *mongo.Collection
	HeaderStorage   *mongo.Collection
	BodyStorage     *mongo.Collection
	ConsumerStorage *mongo.Collection
}

// Scaffold is the main frame for storing references to content in the post
type Scaffold struct {
	ID         primitive.ObjectID   `bson:"_id"`
	OwnedBy    string               `bson:"creator_id"`
	CreatedAt  time.Time            `bson:"created_at"`
	UpdatedAt  time.Time            `bson:"updated_at,omitempty"`
	HeaderRefs []primitive.ObjectID `bson:"header_ids"`
	BodyRefs   []primitive.ObjectID `bson:"body_ids"`
}

// Content is used to store both header and body data in their own collections
type Content struct {
	ID        primitive.ObjectID `bson:"_id"`
	OwnedBy   string             `bson:"creator_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Data      string             `bson:"data,omitempty"`
}

type PostData struct {
	IDHex     string
	Author    string
	CreatedAt time.Time
	Header    string
	Body      string
}

type PostDataHistory struct {
	IDHex     string
	Author    string
	CreatedAt time.Time
	Headers   []string
	Bodies    []string
}
