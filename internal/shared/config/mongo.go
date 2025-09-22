package config

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoResource struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongo(ctx context.Context) (*MongoResource, error) {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_DB")
	if uri == "" || dbName == "" {
		return nil, errors.New("missing MongoDB URI or DB name in environment variables")
	}

	// Timeout de conexion
	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	pingCtx, pingCancel := context.WithTimeout(ctx, 5*time.Second)
	defer pingCancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		return nil, err
	}

	return &MongoResource{
		Client: client,
		DB:     client.Database(dbName),
	}, nil
}

func (m *MongoResource) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
