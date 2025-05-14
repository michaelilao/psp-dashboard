package transaction

import (
	"context"
	"psp-dashboard-be/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collName = "transactions"
var dbName = "psp"

type Store struct {
	client *mongo.Client
}

func NewStore(client *mongo.Client) *Store {
	return &Store{client: client}
}


func (s *Store) CreateTransaction(transaction types.Transaction) (primitive.ObjectID, error){

	coll := s.client.Database(dbName).Collection(collName)
	result, err := coll.InsertOne(context.TODO(), transaction)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
	
}