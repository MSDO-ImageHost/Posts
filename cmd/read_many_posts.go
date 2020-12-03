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
		res.Status.Message = err.Error()
		return res, err
	}

	// Alter database
	storageRes, err := storage.FindManyPosts(postReq.PostIDs)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		res.Status.Message = err.Error()
		return res, err
	}

	// Construct response object
	postResults := make([]api.NoPostHistoryStruct, len(storageRes))
	for i := range storageRes {
		postResults[i] = api.NoPostHistoryStruct{
			PostID:    storageRes[i].IDHex,
			Author:    storageRes[i].Author,
			CreatedAt: storageRes[i].CreatedAt,
			UpdatedAt: storageRes[i].UpdatedAt,
			Header: api.PostContentStruct{
				Author:    storageRes[i].Header.Author,
				Data:      storageRes[i].Header.Data,
				CreatedAt: storageRes[i].Header.CreatedAt,
			},
			Body: api.PostContentStruct{
				Author:    storageRes[i].Body.Author,
				Data:      storageRes[i].Body.Data,
				CreatedAt: storageRes[i].Body.CreatedAt,
			},
		}
	}

	// Parse response object into json
	resBytes, err := json.Marshal(postResults)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		res.Status.Message = err.Error()
		return res, err
	}

	// Set status codes and return
	res.Payload = resBytes
	res.Status.Code = http.StatusCreated
	res.Status.Message = http.StatusText(http.StatusCreated)
	return res, nil
}
