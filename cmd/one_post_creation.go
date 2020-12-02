package main

import (
	"encoding/json"
	"net/http"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

func postCreationHandler(req broker.HandleRequestPayload) (res broker.HandleResponsePayload, err error) {

	// Parse request
	newPost := api.CreateOnePostRequest{}
	if err := json.Unmarshal(req.Payload, &newPost); err != nil {
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = err.Error()
		return res, err
	}

	// Store post into database
	storageRes, err := storage.AddOnePost(storage.PostData{
		Author: req.UserID,
		Header: storage.PostContent{Data: newPost.Header},
		Body:   storage.PostContent{Data: newPost.Body},
	})
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		res.Status.Message = err.Error()
		return res, err
	}

	// Construct response object
	postRes := api.NoHistoryPostResponse{
		PostID:    storageRes.IDHex,
		Author:    storageRes.Author,
		CreatedAt: storageRes.CreatedAt,
		Header: api.PostContent{
			Author:    storageRes.Header.Author,
			Data:      storageRes.Header.Data,
			CreatedAt: storageRes.Header.CreatedAt,
		},
		Body: api.PostContent{
			Author:    storageRes.Body.Author,
			Data:      storageRes.Body.Data,
			CreatedAt: storageRes.Body.CreatedAt,
		},
	}

	// Parse response object into json
	resBytes, err := json.Marshal(postRes)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		res.Status.Message = err.Error()
		return res, err
	}

	res.Payload = resBytes
	res.Status.Code = http.StatusCreated
	res.Status.Message = http.StatusText(http.StatusCreated)
	return res, nil
}
