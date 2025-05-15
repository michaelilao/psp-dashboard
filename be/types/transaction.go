package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TransactionStore interface {
	CreateTransaction(transaction Transaction) (primitive.ObjectID, error)
	GetTransactionsByQuery(filter bson.D) ([]Transaction, error)
	DeleteTransactionById(transactionId primitive.ObjectID) (error)
	UpdateTransactionById(transaction Transaction) (error)
}

type Transaction struct {
	Id    					primitive.ObjectID 	`bson:"_id,omitempty" json:"id"`
	UserId  				primitive.ObjectID  `bson:"userId" json:"userId"`
	Date 						time.Time           `bson:"date" json:"date"`
	Category				string							`bson:"category" json:"category"`
	TransactionType	string							`bson:"transactionType" json:"transactionType"` 
	Amount					int									`bson:"amount" json:"amount"`
	Name						string							`bson:"name" json:"name"`
	Notes 					string						 	`bson:"notes" json:"notes"`
}

type CreateTransactionPayload struct {
	UserId						string 	`json:"userId" validate:"required"`
	Date							string 	`json:"date"`
	Category					string	`json:"category" validate:"required"`
	TransactionType		string 	`json:"transactionType" validate:"required,oneof=income expense"`
	Amount						int  		`json:"amount" validate:"required"`
	Name  						string 	`json:"name"`
	Notes 						string 	`json:"notes"`
}


type UpdateTransactionPayload struct {
	Date							string 	`json:"date"`
	Category					string	`json:"category"`
	TransactionType		string 	`json:"transactionType" validate:"required,oneof=income expense"`
	Amount						int  		`json:"amount" validate:"required"`
	Name  						string 	`json:"name"`
	Notes 						string 	`json:"notes"`
}
