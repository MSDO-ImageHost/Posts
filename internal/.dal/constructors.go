package database

import "github.com/mitchellh/mapstructure"

func NewPost(v interface{}) (post StorageInterface, err error) {
	var postData PostData
	err = mapstructure.Decode(v, &postData)
	return &postData, nil
}
