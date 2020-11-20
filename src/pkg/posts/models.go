package posts

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostQueryID string

// Post is the main frame for storing references to content in the post
type Post struct {
	PostId    primitive.ObjectID `json:"post_id"`
	CreatedAt time.Time          `json:"created_at"`
	AuthorID  string             `json:"author_id"`
	Header    string             `json:"title"`
	Body      string             `json:"body"`
	Meta      Meta               `json:"-"`
}

type Meta struct {
	TZ string
}
