package repository

import (
	"context"
	"time"

	"github.com/BigDwarf/sahtian/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	collection *mongo.Collection
}

func NewClientRepository(db *mongo.Database) *ClientRepository {
	return &ClientRepository{
		collection: db.Collection("clients"),
	}
}

func (r *ClientRepository) Create(ctx context.Context, client *model.Client) error {
	client.ID = primitive.NewObjectID().Hex()
	client.CreatedAt = time.Now().Unix()

	_, err := r.collection.InsertOne(ctx, client)
	return err
}

func (r *ClientRepository) GetAll(ctx context.Context) ([]*model.Client, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var clients []*model.Client
	if err = cursor.All(ctx, &clients); err != nil {
		return nil, err
	}

	return clients, nil
}
