package database

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ParseMongoPostContent(res interface{}) (content PostContent, err error) {

	_map := res.(primitive.A)[0].(primitive.D).Map()
	fmt.Println(_map)
	mapstructure.Decode(_map, &content)

	//content.ID = _map["_id"]
	//content.AuthorID = _map["creator_id"].(string)
	//content.Data = _map["data"].(string)
	content.CreatedAt = _map["created_at"].(primitive.DateTime).Time()

	return content, err
}
