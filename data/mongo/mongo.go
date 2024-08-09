package mongo

import (
	"context"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	db  *mongo.Client
	cfg *config.Config
}

func NewMongo(DB *mongo.Client, cfg *config.Config) *Mongo {
	return &Mongo{
		db:  DB,
		cfg: cfg,
	}
}

func (m *Mongo) Insert(collection string, data interface{}) error {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(collection)

	_, err := collectionConnected.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

//implements all the methods to interact with mongo
