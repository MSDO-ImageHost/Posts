package database

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidHexFormat(hexId string) (err error) {
	_, err = primitive.ObjectIDFromHex(hexId)
	return err
}

func AssertClientInstance() error {
	if shell.Client == nil {
		log.Println("Database:\tNo database client instance! Attempting setup..")
		if err := Init(); err != nil {
			log.Panicln("Database:\tCould not establish connection! Panicking", err)
		}
	}
	return nil
}

func hexes2ObjectID(hexIds []string) (objIds []primitive.ObjectID, err error) {

	// New slice for Object Ids
	objIds = make([]primitive.ObjectID, len(hexIds))

	// Map hex id to object id
	for i, hexId := range hexIds {
		objId, err := primitive.ObjectIDFromHex(hexId)
		if err != nil {
			return nil, err
		}
		objIds[i] = objId
	}
	return objIds, nil
}
