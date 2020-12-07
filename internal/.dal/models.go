package database

import (
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	Client          *mongo.Client
	Db              *mongo.Database
	ScaffoldStorage *mongo.Collection
	HeaderStorage   *mongo.Collection
	BodyStorage     *mongo.Collection
}

// Scaffold is the main frame for storing references to content in the post
type PostData struct {
	ID            interface{} `bson:"_id" mapstructure:"_id"`
	AuthorID      string      `bson:"creator_id" mapstructure:"creator_id"`
	CreatedAt     time.Time   `bson:"created_at" mapstructure:"created_at"`
	UpdatedAt     *time.Time  `bson:"updated_at,omitempty" mapstructure:"updated_at"`
	MarkedDeleted *bool       `bson:"marked_deleted,omitempty" mapstructure:"marked_deleted"`
	Headers       interface{} `bson:"headers" mapstructure:"header"`
	Bodies        interface{} `bson:"bodies" mapstructure:"body"`
}

type PostContent struct {
	ID            interface{} `bson:"_id" mapstructure:"_id"`
	AuthorID      string      `bson:"creator_id" mapstructure:"creator_id"`
	CreatedAt     time.Time   `bson:"created_at" mapstructure:"created_at"`
	Data          string      `bson:"data,omitempty" mapstructure:"data"`
	MarkedDeleted *bool       `bson:"mark_deleted,omitempty" mapstructure:"marked_deleted"`
}
