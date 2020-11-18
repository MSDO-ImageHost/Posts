package main

import (
	"time"

	models "github.com/MSDO-ImageHost/Posts/pkg/posts"
)

type AuthToken string
type PostId string

type PostData struct {
	Title string
	Body  string
}

type Paging struct {
	Start, End int
	Limit      uint
}

type CreateRequest struct {
	AuthToken
	PostData
}

type GetRequest struct {
	AuthToken
	PostId
}

type UpdateRequest struct {
	AuthToken
	PostId
	PostData
}

type DeleteRequest struct {
	AuthToken
	PostId
}

type GetManyRequest struct {
	AuthToken
	PostIds []PostId
	Paging
}

type GetHistoryRequest struct {
	AuthToken
	PostId
	Paging
}

type ResponseWrapper struct {
	Date           time.Time
	Data           models.Post
	StatusCode     int
	Message        string
	ProcessingTime int
	NodeRespondant string
}
