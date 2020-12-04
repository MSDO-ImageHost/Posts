package main

import (
	"encoding/json"
	"net/http"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

func updateOnePostHandler(req broker.HandleRequestPayload) (res broker.HandleResponsePayload, err error) {

	// Parse request
	postReq := api.NoPostHistoryStruct{}
	if err := json.Unmarshal(req.Payload, &postReq); err != nil {
		res.Status.Code = http.StatusBadRequest
		res.Status.Message = err.Error()
		return res, err
	}

	newHeader, updateHeader := postReq.Header.(string)
	newBody, updateBody := postReq.Body.(string)

	// Alter database
	storageRes, err := storage.UpdateOnePost(storage.PostData{
		IDHex: postReq.PostID,
		Header: storage.PostContent{
			Data: newHeader,
			//Author: <nil>,
			Update: updateHeader,
		},
		Body: storage.PostContent{
			Data: newBody,
			//Author: <nil>,
			Update: updateBody,
		},
	})

	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		res.Status.Message = err.Error()
		return res, err
	}

	// Construct response object
	postRes := api.NoPostHistoryStruct{
		PostID:    storageRes.IDHex,
		Author:    storageRes.Author,
		CreatedAt: storageRes.CreatedAt,
		Header: api.PostContentStruct{
			Author:    storageRes.Header.Author,
			Data:      storageRes.Header.Data,
			CreatedAt: storageRes.Header.CreatedAt,
		},
		Body: api.PostContentStruct{
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

	// Set status codes and return
	res.Payload = resBytes
	res.Status.Code = http.StatusCreated
	res.Status.Message = http.StatusText(http.StatusCreated)
	return res, nil
}
