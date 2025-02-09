package mongo

import (
	"context"
	"errors"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"go.mongodb.org/mongo-driver/bson"
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

func (m *Mongo) Close() error {
	err := m.db.Disconnect(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) Insert(ctx context.Context, bookmarks entity.Bookmarks) error {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	_, err := collectionConnected.InsertOne(ctx, bookmarks)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) FindAllByField(ctx context.Context, field string, value string) ([]entity.Bookmarks, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	cursor, err := collectionConnected.Find(ctx, map[string]string{field: value})
	if err != nil {
		return nil, err
	}

	var result []entity.Bookmarks
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (m *Mongo) FindOne(ctx context.Context, id string) (entity.Bookmarks, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	var result entity.Bookmarks //TODO:validar o campo id se é assim mesmo
	err := collectionConnected.FindOne(ctx, map[string]string{"id": id}).Decode(&result)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	return result, nil
}

func (m *Mongo) UpdateOne(ctx context.Context, story *entity.Bookmarks) (entity.Bookmarks, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	result := collectionConnected.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "id", Value: story.ID}},
		bson.M{"$set": story},
	)
	if result == nil {
		return entity.Bookmarks{}, errors.New(constants.ERROR_STORY_NOT_FOUND)
	}

	storyUpdated := entity.Bookmarks{}
	err := result.Decode(&storyUpdated)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	return storyUpdated, nil
}

func (m *Mongo) DeleteOne(ctx context.Context, id string) error {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	_, err := collectionConnected.DeleteOne(ctx, map[string]string{"id": id})
	if err != nil {
		return err
	}

	return nil
}

//implements all the methods to interact with mongo
