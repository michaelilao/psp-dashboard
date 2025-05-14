package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionStore interface {
	CreateTransaction(transaction Transaction) (primitive.ObjectID, error)
}

type Transaction struct {
	ID    					primitive.ObjectID 	`bson:"_id,omitempty"`
	UserID  				primitive.ObjectID  `bson:"userID"`
	Date 						time.Time           `bson:"date"`
	Category				string							`bson:"category"`
	TransactionType	string							`bson:"transactionType"` 
	Amount					int									`bson:"int"`
	Name						string							`bson:"name"`
	Notes 					string						 	`bson:"notes"`
}

type CreateTransactionPayload struct {
	UserID						string 	`json:"userID" validate:"required"`
	Date							string 	`json:"date"`
	Category					string	`json:"category" validate:"required"`
	TransactionType		string 	`json:"TransactionType" validate:"required,oneof=income expense"`
	Amount						int  		`json:"amount" validate:"required"`
	Name  						string 	`json:"name" validate:"required"`
	Notes 						string `json:"notes"`
}
