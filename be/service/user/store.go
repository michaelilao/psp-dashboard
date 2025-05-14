package user

import (
	"context"
	"fmt"
	"psp-dashboard-be/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collName = "users"
var dbName = "psp"

type Store struct {
	client *mongo.Client
}

func NewStore(client *mongo.Client) *Store {
	return &Store{client: client}
}


func (s *Store) DeleteUserById(userId primitive.ObjectID) (error) {
	coll := s.client.Database(dbName).Collection(collName)
	
	filter := bson.D{{Key: "_id", Value: userId}}
	result, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0{
		return fmt.Errorf("userid does not exist")
	}
	return nil
}


func (s *Store) UpdateUserById(user types.User) (error) {

	coll := s.client.Database(dbName).Collection(collName)
	filter := bson.D{{Key: "_id", Value: user.Id}}
	
	result, err := coll.ReplaceOne(context.TODO(), filter, user)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("userid does not exist")
	}

	return nil
}


func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	coll := s.client.Database(dbName).Collection(collName)
	
	user := &types.User{}
	err := coll.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)

	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (s *Store) GetUserById(id string) (*types.User, error) {
	coll := s.client.Database(dbName).Collection(collName)
	user := &types.User{}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = coll.FindOne(context.TODO(), bson.D{{Key: "_id", Value: objectId}}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}



func (s *Store) InsertUser(user types.User) (primitive.ObjectID, error) {

	coll := s.client.Database(dbName).Collection(collName)
	result, err := coll.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (s *Store) GetUsers() ([]types.User, error) {
	coll := s.client.Database(dbName).Collection(collName)

	var users []types.User
	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		return users, err
	}

	err = cursor.All(context.TODO(), &users); 
	if err != nil {
		return users, err
	}
	
	return users, nil
}