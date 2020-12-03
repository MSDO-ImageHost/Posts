package api

import (
	"time"
)

type SinglePostID struct {
	PostID string `json:"post_id"`
}

type PostContentStruct struct {
	Author    string     `json:"author_id"`
	CreatedAt *time.Time `json:"created_at"`
	Data      string     `json:"data"`
}

type StatusCode struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

type PagingStruct struct {
	Start interface{} `json:"start"`
	End   interface{} `json:"end"`
	Limit uint        `json:"limit"`
}

type NoPostHistoryStruct struct {
	PostID    string        `json:"post_id"`
	Author    string        `json:"author_id"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
	Header    interface{}   `json:"header"` // Used as string for unmarshalling and PostContentStruct for marshalling
	Body      interface{}   `json:"body"`
	Paging    *PagingStruct `json:"paging,omitempty"`
}

type PostHistoryStruct struct {
	PostID    string              `json:"post_id"`
	Author    string              `json:"author_id"`
	CreatedAt *time.Time          `json:"created_at"`
	UpdatedAt *time.Time          `json:"updated_at,omitempty"`
	Headers   []PostContentStruct `json:"header"`
	Bodies    []PostContentStruct `json:"body"`
	Paging    *PagingStruct       `json:"paging,omitempty"`
}

type ManyPostIds struct {
	PostIDs []string      `json:"post_ids"`
	Paging  *PagingStruct `json:"paging,omitempty"`
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
