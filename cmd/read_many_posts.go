package main

import (
	"encoding/json"
	"net/http"

	api "github.com/MSDO-ImageHost/Posts/internal/api"
	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

func readManyPostsHandler(req broker.HandleRequestPayload) (res broker.HandleResponsePayload, err error) {

	// Parse request
	postReq := api.ManyPostIds{}
	if err := json.Unmarshal(req.Payload, &postReq); err != nil {
		res.Status.Code = http.StatusBadRequest
		return res, err
	}

	// Alter database
	storageRes, err := storage.FindManyPosts(postReq.PostIDs, *postReq.Paging)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		return res, err
	}

	// Construct response object
	postResults := make([]api.NoPostHistoryStruct, len(storageRes))
	for i := range storageRes {
		postResults[i] = api.NoPostHistoryStruct{
			PostID:    storageRes[i].IDHex,
			AuthorID:  storageRes[i].AuthorID,
			CreatedAt: storageRes[i].CreatedAt,
			UpdatedAt: storageRes[i].UpdatedAt,
			Header: api.PostContentStruct{
				AuthorID:  storageRes[i].Header.AuthorID,
				Data:      storageRes[i].Header.Data,
				CreatedAt: storageRes[i].Header.CreatedAt,
			},
			Body: api.PostContentStruct{
				AuthorID:  storageRes[i].Body.AuthorID,
				Data:      storageRes[i].Body.Data,
				CreatedAt: storageRes[i].Body.CreatedAt,
			},
		}
	}

	// Parse response object into json
	resBytes, err := json.Marshal(postResults)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		return res, err
	}

	// Set status codes and return
	res.Payload = resBytes
	res.Status.Code = http.StatusCreated
	return res, nil
}
