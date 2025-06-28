package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	context  context.Context
	client   *mongo.Client
	database *mongo.Database

	Collections struct {
		Logins          *mongo.Collection
		Recommendations *mongo.Collection
	}
}
