package main

import (
	"encoding/json"
	"net/http"

	"github.com/MSDO-ImageHost/Posts/internal/api"
	"github.com/MSDO-ImageHost/Posts/internal/auth"
	broker "github.com/MSDO-ImageHost/Posts/internal/broker"
	storage "github.com/MSDO-ImageHost/Posts/internal/database"
)

func deleteManyPostsHandler(req broker.HandleRequestPayload) (res broker.HandleResponsePayload, err error) {

	headers, err := api.ParseHeader(req.Headers)
	if err != nil {
		return res, err
	}

	userAuth, err := auth.AuthJWT(headers.JWT)
	if err != nil {
		res.Status.Code = http.StatusUnauthorized
		return res, err
	}

	// Parse request
	postReq := api.ManyPostIds{}
	if err := json.Unmarshal(req.Payload, &postReq); err != nil {
		res.Status.Code = http.StatusBadRequest
		return res, err
	}

	// Alter database
	storageRes, err := storage.DeleteManyPosts(postReq.PostIDs, userAuth)
	if err != nil {
		res.Status.Code = http.StatusInternalServerError
		return res, err
	}

	// Construct response object
	postRes := api.ManyPostIds{PostIDs: storageRes}

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
