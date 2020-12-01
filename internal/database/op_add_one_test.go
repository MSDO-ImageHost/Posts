package database_test

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/MSDO-ImageHost/Posts/internal/database"
)

func TestAdd(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	s := database.NewMockstorageInterface(mockCtrl)

	post := database.PostData{
		Author: "123-christian-id",
		Header: database.PostContent{Data: "Hello from header!"},
		Body:   database.PostContent{Data: "Hello from body!"},
	}

	s.EXPECT().AddOne(post).Return(post, errors.New("DB error"))
}
