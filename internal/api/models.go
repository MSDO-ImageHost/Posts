package api

import (
	"time"
)

type CreateRequest struct {
	AuthToken string `json:"auth_token"`
	Header    string `json:"header"`
	Body      string `json:"body"`
}

type CreateResponse struct {
	PostID string `json:"post_id"`
}

type GetRequest struct {
	AuthToken string `json:"auth_token"`
	PostID    string `json:"post_id"`
}
type GetResponse struct {
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	AuthorID  string    `json:"author_id"`
	Header    string    `json:"header"`
	Body      string    `json:"body"`
}

type UpdateRequest struct {
	AuthToken string `json:"auth_token"`
	PostID    string `json:"post_id"`
	Header    string `json:"header"`
	Body      string `json:"body"`
}
type UpdateResponse struct {
	PostID string `json:"post_id"`
}

type DeleteRequest struct {
	AuthToken string `json:"auth_token"`
	PostID    string `json:"post_id"`
}
type DeleteResponse struct {
	PostID string `json:"post_id"`
}

type GetManyRequest struct {
	AuthToken string   `json:"auth_token"`
	PostIDS   []string `json:"post_ids"`
	Paging    `json:"paging"`
}

type GetManyResponse struct {
	Posts []GetResponse
}

type GetUserPostsResponse GetManyResponse

type GetHistoryRequest struct {
	AuthToken string `json:"auth_token"`
	PostID    string `json:"post_id"`
	Paging    `json:"paging"`
}
type GetHistoryResponse struct {
}

type ContentHistory struct {
	CreatedAt time.Time `json:"created_at"`
	Header    string    `json:"header"`
	Body      string    `json:"body"`
}

type Paging struct {
	Start interface{} `json:"start"`
	End   interface{} `json:"end"`
	Limit uint        `json:"limit"`
}

type ResponseWrapper struct {
	Data             interface{} `json:"data"`
	ProcessingTimeNs int64       `json:"processing_time_ns"`
	StatusCode       uint        `json:"-"`
	Message          string      `json:"-"`
	NodeRespondant   interface{} `json:"-"`
}
