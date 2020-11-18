package posts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Field types
type PostQueryID string
type Author string
type _ID primitive.ObjectID
type CreatedAt time.Time
type ReplacedBy primitive.ObjectID

// Content store field specific data for a post. It also contain references to revisions of itself
type Content struct {
	_ID        _ID        `bson:"_id" json:"-"`
	Author     Author     `bson:"author,omitempty" json:"-"`
	Data       string     `bson:"data,omitempty" json:"-"`
	ReplacedBy ReplacedBy `bson.ObjectId:"title,omitempty" json:"-"`
}

// Post is the main frame for storing references to content in the post
type Post struct {
	_ID       _ID       `bson:"_id" json:"post_id"`
	CreatedAt CreatedAt `json:"created_at"`
	Author    Author    `bson:"author,omitempty" json:"creator_id"`
	Title     Content   `bson.ObjectId:"title,omitempty" json:"title"`
	Body      Content   `bson.ObjectId:"body,omitempty" json:"body"`
}
