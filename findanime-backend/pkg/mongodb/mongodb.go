package mongodb

import (
	"api/internal/config"
	"context"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// DB is the exported database connection instance
	DB = New()
)

func New() *MongoDB {
	var db MongoDB

	mongoOptions := options.Client().ApplyURI(config.Configuration.DatabaseURI)

	// Create a new context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, mongoOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return nil
	}

	// Assign the connected client to the db struct
	db.client = client
	db.context = ctx

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return nil
	}

	db.database = db.client.Database("findanime")
	db.Collections.Recommendations = db.database.Collection("recommendations")
	return &db
}

func (c *MongoDB) CountCollection(filter bson.M, collection *mongo.Collection) (int64, error) {
	cursor, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return cursor, nil
}

func (c *MongoDB) Find(queryData bson.M, collection *mongo.Collection) (*mongo.Cursor, error) {
	cursor, err := collection.Find(context.Background(), queryData)
	if err != nil {
		return nil, err
	}

	return cursor, nil
}

func (c *MongoDB) FindOne(queryData bson.M, collection *mongo.Collection, useSort ...bool) (bson.M, error) {
	var reply bson.M

	if len(useSort) != 0 {
		option := options.FindOneOptions{
			Sort: bson.D{
				{Key: "_id", Value: -1},
			},
		}
		err := collection.FindOne(context.Background(), queryData, &option).Decode(&reply)
		if err != nil {
			return nil, err
		}

		return reply, nil
	}

	err := collection.FindOne(context.Background(), queryData).Decode(&reply)
	if err != nil {
		return nil, err
	}

	return reply, nil

}

func (c *MongoDB) FindOneAndUpdate(queryData bson.M, update bson.M, collection *mongo.Collection, useSort ...bool) (bson.M, error) {
	var reply bson.M

	if len(useSort) != 0 {
		option := options.FindOneAndUpdateOptions{
			Sort: bson.D{
				{Key: "_id", Value: -1},
			},
		}
		err := collection.FindOneAndUpdate(context.Background(), queryData, update, &option).Decode(&reply)
		if err != nil {
			return nil, err
		}

		return reply, nil
	}

	err := collection.FindOneAndUpdate(context.Background(), queryData, update).Decode(&reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (c *MongoDB) InsertOne(queryData bson.M, collection *mongo.Collection) (*mongo.InsertOneResult, error) {
	reply, err := collection.InsertOne(context.Background(), queryData)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (c *MongoDB) DeleteOne(queryData bson.M, collection *mongo.Collection) (*mongo.DeleteResult, error) {
	reply, err := collection.DeleteOne(context.Background(), queryData)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
