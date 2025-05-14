package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	InsertUser(user User) (primitive.ObjectID, error)
	GetUsers() ([]User, error)
}

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
	Notes string						 `bson:"notes"`
}

type RegisterUserPayload struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Notes string `json:"notes"`
}
