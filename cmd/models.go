package main


import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post is the main frame for storing references to content in the post
type Post struct {
	ID 			primitive.ObjectID 	`bson:"_id" json:"post_id"`
	CreatedAt 	time.Time         	`bson:"created_at" json:"created_at"`
	Author 		string             	`bson:"author,omitempty" json:"creator_id"`
	Title 		Content				`bson.ObjectId:"title,omitempty" json:"title"`
	Body 		Content				`bson.ObjectId:"body,omitempty" json:"body"`
}

// Content store field specific data for a post. It also contain references to revisions of itself
type Content struct {
	ID 			primitive.ObjectID 	`bson:"_id" json:"-"`
	Author 		string             	`bson:"author,omitempty" json:"-"`
	CreatedAt 	time.Time         	`bson:"created_at" json:"-"`
	Data 		string				`bson:"data,omitempty" json:"-"`
	ReplacedBy 	primitive.ObjectID 	`bson.ObjectId:"title,omitempty" json:"-"`
}

