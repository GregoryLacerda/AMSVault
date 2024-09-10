package mongo

import (
	"context"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/entity"
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

func (m *Mongo) FindAllByField(collection string, field string, value string) ([]entity.Story, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(collection)

	cursor, err := collectionConnected.Find(context.TODO(), map[string]string{field: value})
	if err != nil {
		return nil, err
	}

	var result []entity.Story
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *Mongo) FindByID(collection string, id string) (entity.Story, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(collection)

	var result entity.Story
	err := collectionConnected.FindOne(context.TODO(), map[string]string{"id": id}).Decode(&result)
	if err != nil {
		return entity.Story{}, err
	}

	return result, nil
}

func (m *Mongo) Delete(collection string, id string) error {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(collection)

	_, err := collectionConnected.DeleteOne(context.TODO(), map[string]string{"id": id})
	if err != nil {
		return err
	}

	return nil
}

//implements all the methods to interact with mongo
