package db

import (
	"context"
	"errors"
	"log"
	"test-hex-architecture/internal/shared/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoResource struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongo(ctx context.Context) (*MongoResource, error) {
	user, errUser := config.MongoUser()
	pass, errPass := config.MongoPass()
	host, errHost := config.MongoHost()
	port, errPort := config.MongoPort()
	dbName, errDBName := config.MongoDBName()

	if errUser != nil || errPass != nil || errHost != nil || errPort != nil || errDBName != nil {
		return nil, errors.New("error retrieving MongoDB configuration from environment variables")
	}

	uri := "mongodb://" + user + ":" + pass + "@" + host + ":" + port
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

	log.Printf("Connected to MongoDB at %s, using database: %s", host, dbName)

	return &MongoResource{
		Client: client,
		DB:     client.Database(dbName),
	}, nil
}

func (m *MongoResource) Disconnect(ctx context.Context) error {
	return m.Client.Disconnect(ctx)
}
