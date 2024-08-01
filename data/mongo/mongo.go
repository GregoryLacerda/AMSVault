package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	db *mongo.Client
}

func NewMongo(DB *mongo.Client) *Mongo {
	return &Mongo{
		db: DB,
	}
}

//implements all the methods to interact with mongo
