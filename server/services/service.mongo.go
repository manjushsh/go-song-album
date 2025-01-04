package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName          = "go_projects"
	connectTimeout  = 10 * time.Second
	operationTimeout = 5 * time.Second
	findAllTimeout  = 30 * time.Second
)

type MongoService struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoService() (*MongoService, error) {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		return nil, fmt.Errorf("MONGO_URI environment variable not set")
	}

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return &MongoService{
		client: client,
		db:     client.Database(dbName),
	}, nil
}

func (s *MongoService) Insert(collection string, document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	return s.db.Collection(collection).InsertOne(ctx, document)
}

func (s *MongoService) InsertMany(collection string, documents []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	return s.db.Collection(collection).InsertMany(ctx, documents)
}

func (s *MongoService) FindOne(collection string, filter interface{}) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	return s.db.Collection(collection).FindOne(ctx, filter)
}

func (s *MongoService) FindAll(collection string, filter interface{}) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), findAllTimeout)
	defer cancel()

	return s.db.Collection(collection).Find(ctx, filter)
}

func (s *MongoService) Update(collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	return s.db.Collection(collection).UpdateOne(ctx, filter, update)
}

func (s *MongoService) Delete(collection string, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), operationTimeout)
	defer cancel()

	return s.db.Collection(collection).DeleteOne(ctx, filter)
}

func (s *MongoService) Disconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	return s.client.Disconnect(ctx)
}
