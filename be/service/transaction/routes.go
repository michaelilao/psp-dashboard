package transaction

import (
	"fmt"
	"log"
	"net/http"
	"psp-dashboard-be/types"
	"psp-dashboard-be/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	store types.TransactionStore
	userStore types.UserStore
}


func NewHandler(store types.TransactionStore, userStore types.UserStore) *Handler{
	return &Handler{
		store: store,
		userStore: userStore,
	}
}


func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /transaction", h.HandleCreateTransaction)
	router.HandleFunc("GET /transaction", h.HandleGetTransactions)
	router.HandleFunc("DELETE /transaction/{transactionID}", h.HandleDeleteTransactionByID)
	router.HandleFunc("PUT /transaction/{transactionID}", h.HandleUpdateTransactionByID)
}


func (h *Handler) HandleUpdateTransactionByID(w http.ResponseWriter, r *http.Request) {
		transactionID := r.PathValue("transactionID")
		if transactionID == "" {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("provide non empty transactionid"))
			return
		}

		objectId, err := primitive.ObjectIDFromHex(transactionID)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid transactionid %v", err))
			return
		}

		var payload types.UpdateTransactionPayload
		if err := utils.ParseJSON(r, &payload); err != nil {
			utils.WriteError(w, http.StatusBadRequest, err)
			return
		}

		if err := utils.Validate.Struct(payload); err != nil {
			errors := err.(validator.ValidationErrors)
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
			return
		}

		// Check if transaction exists
		query := bson.D{{Key: "_id", Value: objectId}}
		transaction, err := h.store.GetTransactionsByQuery(query)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error fetching transaction %v", err))
			return
		}

		if len(transaction) == 0 {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("transaction does not exist %v", err))
			return
		}

		existingTransaction := transaction[0]

		existingTransaction.Category = payload.Category
		existingTransaction.TransactionType = payload.TransactionType
		existingTransaction.Name = payload.Name
		existingTransaction.Notes = payload.Notes
		existingTransaction.Amount = payload.Amount

		date, err := time.Parse("2006-01-02", payload.Date)
		if err == nil {
			existingTransaction.Date = date
		}

		err = h.store.UpdateTransactionByID(existingTransaction)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating transaction %v", err))
			return
		}



		utils.WriteJSON(w, http.StatusOK, existingTransaction)
		
}


func (h *Handler) HandleDeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
		transactionID := r.PathValue("transactionID")
		if transactionID == "" {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("provide non empty transactionid"))
			return
		}

		objectId, err := primitive.ObjectIDFromHex(transactionID)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid transactionid %v", err))
			return
		}

		err = h.store.DeleteTransactionByID(objectId)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error deleting transaction %v", err))
			return
		}
	
		utils.WriteJSON(w, http.StatusOK, true)

}

func (h *Handler) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	  // Possible Queries

		// can take as strings
		query, err := CreateQuery(r.URL.Query())
		if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid query %v", err))
				return
		}
			
		log.Println(query)
		transactions, err := h.store.GetTransactionsByQuery(query)				
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error fetching transactions %v", err))
			return
		}

	
		utils.WriteJSON(w, http.StatusOK, transactions)
}

func (h *Handler) HandleCreateTransaction(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateTransactionPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	// Check if user exists
	user, err := h.userStore.GetUserByID(payload.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user does not exist %v", err))
		return
	}

	
	newTransaction := types.Transaction{
		UserID: 					user.ID,	
		Category: 				payload.Category,
		TransactionType: 	payload.TransactionType,
		Amount: 					payload.Amount,
		Name: 						payload.Name,
		Notes: 						payload.Notes,
	}

	// yyyy-mm-ddd
	date, err := time.Parse("2006-01-02", payload.Date)
	if err != nil {
		date = time.Now()
	}
	newTransaction.Date = date


	id, err := h.store.CreateTransaction(newTransaction)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error inserting new transaction %v", err))
		return
	}

	newTransaction.ID = id
	utils.WriteJSON(w, http.StatusOK, newTransaction)
}