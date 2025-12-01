package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnector struct {
	cfg    Config
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoConnector(cfg Config) *MongoConnector {
	return &MongoConnector{cfg: cfg}
}

func (mc *MongoConnector) Connect(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mc.cfg.Uri))
	if err != nil {
		return err
	}
	mc.client = client
	mc.db = client.Database(mc.cfg.Database)
	return nil
}

func (mc *MongoConnector) Disconnect(ctx context.Context) error {
	if mc.client != nil {
		return mc.client.Disconnect(ctx)
	}
	return nil
}

func (mc *MongoConnector) Ping(ctx context.Context) error {
	return mc.client.Ping(ctx, nil)
}

func (mc *MongoConnector) Client() any {
	return mc.client
}

func (mc *MongoConnector) DB() *mongo.Database {
	return mc.db
}

func (mc *MongoConnector) Collection(name string) *mongo.Collection {
	return mc.db.Collection(name)
}
