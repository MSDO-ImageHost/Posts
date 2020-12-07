package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (pd *PostData) Insert() error {

	// Validate that header and body data is present
	headerString, headerOk := pd.Headers.(string)
	bodyString, bodyOk := pd.Bodies.(string)
	if !headerOk || !bodyOk {
		return fmt.Errorf("Missing post content")
	}

	now := time.Now()

	// Construct post content components
	header := PostContent{
		ID:        primitive.NewObjectID(),
		AuthorID:  pd.AuthorID,
		Data:      headerString,
		CreatedAt: now,
	}

	body := PostContent{
		ID:        primitive.NewObjectID(),
		AuthorID:  pd.AuthorID,
		Data:      bodyString,
		CreatedAt: now,
	}

	scaffold := PostData{
		ID:        primitive.NewObjectID(),
		AuthorID:  pd.AuthorID,
		CreatedAt: now,
		Headers:   []primitive.ObjectID{header.ID.(primitive.ObjectID)},
		Bodies:    []primitive.ObjectID{body.ID.(primitive.ObjectID)},
	}

	// Insert components into their respective collections
	_, err := db.HeaderStorage.InsertOne(context.TODO(), header)
	if err != nil {
		return err
	}

	_, err = db.BodyStorage.InsertOne(context.TODO(), body)
	if err != nil {
		return err
	}

	_, err = db.ScaffoldStorage.InsertOne(context.TODO(), scaffold)
	if err != nil {
		return err
	}

	// Populate post data with new info
	pd.ID = scaffold.ID.(primitive.ObjectID).Hex()
	pd.CreatedAt = scaffold.CreatedAt
	pd.Headers = header
	pd.Bodies = body
	return nil
}
