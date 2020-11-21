package posts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostQueryID string

// Scaffold is the main frame for storing references to content in the post
type Post struct {
	ID        interface{} `bson:"_id" json:"post_id"`
	CreatedAt time.Time   `bson:"created_at" json:"created_at"`
	AuthorID  string      `bson:"author_id" json:"author_id"`
	Header    interface{} `bson:"header_ids" json:"title"`
	Body      interface{} `bson:"body_ids" json:"body"`
}

type Content struct {
	ID        primitive.ObjectID `bson:"_id"`
	AuthorID  string             `bson:"author_id"`
	CreatedAt time.Time          `bson:"created_at"`
	Data      string             `bson:"data"`
	Replacing primitive.ObjectID `bson:"replacing,omitempty"`
}
