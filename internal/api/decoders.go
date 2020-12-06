package api

import "github.com/mitchellh/mapstructure"

func ParseHeader(_headers map[string]interface{}) (headers Headers, err error) {
	err = mapstructure.Decode(_headers, &headers)
	return headers, err
}
