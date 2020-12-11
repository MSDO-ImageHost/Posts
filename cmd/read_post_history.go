package main

import (
	"encoding/json"
	"net/http"

	api "github.com/MSDO-ImageHost/Posts/internal/api"
	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

func readPostHistoryHandler(req broker.HandleRequestPayload) (res broker.HandleResponsePayload, err error) {

	// Parse request
	postReq := api.SinglePostID{}
	if err := json.Unmarshal(req.Payload, &postReq); err != nil {
		res.Status.Code = http.StatusBadRequest
		return res, err
	}

	// Query database
	storageRes, err := storage.FindPostHistory(postReq.PostID)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		return res, err
	}

	// Map headers to API struct
	headers := make([]api.PostContentStruct, len(storageRes.Headers))
	for i := range storageRes.Headers {
		headers[i] = api.PostContentStruct{
			AuthorID:  storageRes.Headers[i].AuthorID,
			CreatedAt: storageRes.Headers[i].CreatedAt,
			Data:      storageRes.Headers[i].Data,
		}
	}

	// Map bodies to API struct
	bodies := make([]api.PostContentStruct, len(storageRes.Bodies))
	for i := range storageRes.Bodies {
		bodies[i] = api.PostContentStruct{
			AuthorID:  storageRes.Bodies[i].AuthorID,
			CreatedAt: storageRes.Bodies[i].CreatedAt,
			Data:      storageRes.Bodies[i].Data,
		}
	}

	// Construct response object
	postRes := api.PostHistoryStruct{
		PostID:    storageRes.IDHex,
		AuthorID:  storageRes.AuthorID,
		CreatedAt: storageRes.CreatedAt,
		UpdatedAt: storageRes.UpdatedAt,
		Headers:   headers,
		Bodies:    bodies,
	}

	// Parse response object into json
	resBytes, err := json.Marshal(postRes)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		return res, err
	}

	// Set status codes and return
	res.Payload = resBytes
	res.Status.Code = http.StatusCreated
	return res, nil
}
