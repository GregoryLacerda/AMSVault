package mongo

import (
	"context"
	"errors"
	"time"

	"github.com.br/GregoryLacerda/AMSVault/config"
	"github.com.br/GregoryLacerda/AMSVault/constants"
	"github.com.br/GregoryLacerda/AMSVault/data/model"
	"github.com.br/GregoryLacerda/AMSVault/entity"
	"github.com/google/uuid"
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

func (m *Mongo) Insert(ctx context.Context, userID int64, storyID int64) error {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	location, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return err
	}

	bookmarks := model.Bookmarks{
		ID:        uuid.NewString(),
		UserID:    userID,
		StoryID:   storyID,
		CreatedAt: time.Now().In(location),
		DeletedAt: time.Date(01, 01, 01, 00, 00, 00, 00, location),
		UpdatedAt: time.Date(01, 01, 01, 00, 00, 00, 00, location),
	}

	_, err = collectionConnected.InsertOne(ctx, bookmarks)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) FindAllByUser(ctx context.Context, userID int64) (retVal []entity.Bookmarks, err error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	cursor, err := collectionConnected.Find(ctx, map[string]int64{"user_id": userID})
	if err != nil {
		return nil, err
	}

	var result []model.Bookmarks
	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	for _, bookmark := range result {
		retVal = append(retVal, bookmark.ToEntity())
	}

	return retVal, nil
}

func (m *Mongo) FindOne(ctx context.Context, id string) (entity.Bookmarks, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	var result model.Bookmarks
	err := collectionConnected.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	return result.ToEntity(), nil
}

func (m *Mongo) UpdateOne(ctx context.Context, bookmarks *entity.Bookmarks) (entity.Bookmarks, error) {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	result := collectionConnected.FindOneAndUpdate(
		ctx,
		bson.D{{Key: "_id", Value: bookmarks.ID}},
		bson.M{"$set": bookmarks},
	)
	if result == nil {
		return entity.Bookmarks{}, errors.New(constants.ERROR_STORY_NOT_FOUND)
	}

	bookmarksUpdated := model.Bookmarks{}
	err := result.Decode(&bookmarksUpdated)
	if err != nil {
		return entity.Bookmarks{}, err
	}

	return bookmarksUpdated.ToEntity(), nil
}

func (m *Mongo) DeleteOne(ctx context.Context, id string) error {
	collectionConnected := m.db.Database(m.cfg.MongoDB).Collection(m.cfg.MongoCollection)

	_, err := collectionConnected.DeleteOne(ctx, map[string]string{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
