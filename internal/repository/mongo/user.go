package mongo

import (
	"WebApp/internal/core"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

//func (repository *UserRepository) GetAll(ctx context.Context) ([]*core.User, error) {
//	return nil
//}

func (repository *UserRepository) GetById(ctx context.Context, id string) (*core.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}

	user := &core.User{}

	err = repository.collection.FindOne(ctx, filter).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil

}
