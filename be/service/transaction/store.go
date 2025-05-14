package transaction

import (
	"context"
	"fmt"
	"psp-dashboard-be/types"

	"go.mongodb.org/mongo-driver/bson"
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

func (s *Store) GetTransactionsByQuery(filter bson.D) ([]types.Transaction, error) {
	coll := s.client.Database(dbName).Collection(collName)

	var transactions []types.Transaction
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return transactions, err
	}

	err = cursor.All(context.TODO(), &transactions); 
	if err != nil {
		return transactions, err
	}
	
	return transactions, nil
}

func (s* Store) DeleteTransactionByID(trainsactionID primitive.ObjectID) (error) {
	coll := s.client.Database(dbName).Collection(collName)
	
	filter := bson.D{{Key: "_id", Value: trainsactionID}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0{
		return fmt.Errorf("transactionid does not exist")
	}
	return nil
}