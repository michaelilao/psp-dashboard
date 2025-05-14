package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id string) (*User, error)
	InsertUser(user User) (primitive.ObjectID, error)
	DeleteUserById(primitive.ObjectID) (error)
	UpdateUserById(user User) (error)
 	GetUsers() ([]User, error)

}

type User struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Notes string						 `bson:"notes"`
}

type CreateUserPayload struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Notes string `json:"notes"`
}
