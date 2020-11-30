package database_test

import (
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/MSDO-ImageHost/Posts/internal/database"
)

func TestAdd(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	storage := database.NewMockstorageInterface(mockCtrl)
	storage.Add(database.Scaffold{})
}
