package api

import (
	"time"
)

type SinglePostID struct {
	PostID string `json:"post_id"`
}

type PostContentStruct struct {
	AuthorID  string     `json:"author_id"`
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
	AuthorID  string        `json:"author_id"`
	CreatedAt *time.Time    `json:"created_at"`
	UpdatedAt *time.Time    `json:"updated_at,omitempty"`
	Header    interface{}   `json:"header"` // Used as string for unmarshalling and PostContentStruct for marshalling
	Body      interface{}   `json:"body"`   // Used as string for unmarshalling and PostContentStruct for marshalling
	ImageData *[]byte       `json:"image_data,omitempty"`
	Paging    *PagingStruct `json:"paging,omitempty"`
}

type PostHistoryStruct struct {
	PostID    string              `json:"post_id"`
	AuthorID  string              `json:"author_id"`
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
