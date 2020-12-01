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
type mongoScaffoldRefs struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Author        string               `bson:"creator_id"`
	CreatedAt     time.Time            `bson:"created_at"`
	UpdatedAt     time.Time            `bson:"updated_at,omitempty"`
	MarkedDeleted bool                 `bson:"marked_deleted,omitempty"`
	HeaderRefs    []primitive.ObjectID `bson:"headers"`
	BodyRefs      []primitive.ObjectID `bson:"bodies"`
}

type mongoScaffoldContents struct {
	ID             primitive.ObjectID `bson:"_id"`
	Author         string             `bson:"creator_id"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at,omitempty"`
	MarkedDeleted  bool               `bson:"marked_deleted,omitempty"`
	HeaderContents []mongoContent     `bson:"headers"`
	BodyContents   []mongoContent     `bson:"bodies"`
}

// Content is used to store both header and body data in their own collections
type mongoContent struct {
	ID          primitive.ObjectID `bson:"_id"`
	Author      string             `bson:"creator_id"`
	CreatedAt   time.Time          `bson:"created_at"`
	Data        string             `bson:"data,omitempty"`
	MarkDeleted bool               `bson:"mark_deleted,omitempty"`
}

type PostContent struct {
	Author    string
	CreatedAt time.Time
	Data      string
	Update    bool
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
	UpdatedAt time.Time
	Headers   []PostContent
	Bodies    []PostContent
}
