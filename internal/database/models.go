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
type mongoScaffold struct {
	ID          primitive.ObjectID   `bson:"_id"`
	IDHex       string               `bson:"-"`
	Author      string               `bson:"creator_id"`
	CreatedAt   time.Time            `bson:"created_at"`
	UpdatedAt   time.Time            `bson:"updated_at,omitempty"`
	MarkDeleted bool                 `bson:"mark_deleted,omitempty"`
	HeaderRefs  []primitive.ObjectID `bson:"header_ids"`
	BodyRefs    []primitive.ObjectID `bson:"body_ids"`
	Header      mongoContent         `bson:"header,omitempty"`
	Body        mongoContent         `bson:"body,omitempty"`
	Headers     []mongoContent       `bson:"headers,omitempty"`
	Bodies      []mongoContent       `bson:"bodies,omitempty"`
}

// Content is used to store both header and body data in their own collections
type mongoContent struct {
	ID          primitive.ObjectID `bson:"_id"`
	IDHex       string             `bson:"-"`
	Author      string             `bson:"creator_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	Data        string             `bson:"data,omitempty"`
	MarkDeleted bool               `bson:"mark_deleted,omitempty"`
}

type PostContent struct {
	Author    string
	CreatedAt time.Time
	Data      string
}
type PostData struct {
	IDHex     string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Header    PostContent
	Body      PostContent
}
type PostDataHistory struct {
	IDHex     string
	Author    string
	CreatedAt time.Time
	Headers   []PostContent
	Bodies    []PostContent
}
