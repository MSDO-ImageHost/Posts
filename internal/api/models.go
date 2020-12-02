package api

import (
	"time"
)

// Requests
type CreateOnePostRequest struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}

// Responses
type NoHistoryPostResponse struct {
	PostID    string      `json:"post_id"`
	Author    string      `json:"author_id"`
	CreatedAt *time.Time  `json:"created_at"`
	UpdatedAt *time.Time  `json:"updated_at,omitempty"`
	Header    PostContent `json:"header"`
	Body      PostContent `json:"body"`
}

// Structures
type PostContent struct {
	Author    string     `json:"author_id"`
	CreatedAt *time.Time `json:"created_at"`
	Data      string     `json:"data"`
}

type StatusCode struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

type Paging struct {
	Start interface{} `json:"start"`
	End   interface{} `json:"end"`
	Limit uint        `json:"limit"`
}

type ResponseWrapper struct {
	Data             interface{} `json:"data"`
	ProcessingTimeNs int64       `json:"processing_time_ns"`
	NodeRespondant   interface{} `json:"-"`
}

/**
type UpdateRequest struct {
	PostID string `json:"post_id"`
	Header string `json:"header"`
	Body   string `json:"body"`
}

type GetRequest struct {
	PostID string `json:"post_id"`
}

type DeleteRequest struct {
	PostID string `json:"post_id"`
}
type GetManyRequest struct {
	PostIDS []string `json:"post_ids"`
	Paging  Paging   `json:"paging"`
}




type CreateOnePostResponse struct {
	PostID string `json:"post_id"`
}

type GetResponse struct {
	PostID    string    `json:"post_id"`
	CreatedAt *time.Time `json:"created_at"`
	AuthorID  string    `json:"author_id"`
	Header    string    `json:"header"`
	Body      string    `json:"body"`
}

type UpdateResponse struct {
	PostID string `json:"post_id"`
}

type DeleteResponse struct {
	PostID string `json:"post_id"`
}

type GetManyResponse struct {
	Posts []GetResponse
}

type GetUserPostsResponse GetManyResponse

type GetHistoryRequest struct {
	PostID string `json:"post_id"`
	Paging `json:"paging"`
}
type GetHistoryResponse struct {
}

type ContentHistory struct {
	CreatedAt *time.Time `json:"created_at"`
	Header    string    `json:"header"`
	Body      string    `json:"body"`
}
**/
