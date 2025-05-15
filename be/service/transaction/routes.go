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
	router.HandleFunc("DELETE /transaction/{transactionId}", h.HandleDeleteTransactionById)
	router.HandleFunc("PUT /transaction/{transactionId}", h.HandleUpdateTransactionById)
}


// @Summary Updates Transaction by Id
// @tags Transaction
// @Produce json
// @Accept  json
// @Param   transaction  body     types.UpdateTransactionPayload  true  "Transaction to update"
// @Param id path string true "Transaction ID"
// @Success 200 {object} types.Transaction
// @Router /transaction/{transactionId} [PUT]
func (h *Handler) HandleUpdateTransactionById(w http.ResponseWriter, r *http.Request) {
		transactionId := r.PathValue("transactionId")
		if transactionId == "" {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("provide non empty transactionid"))
			return
		}

		objectId, err := primitive.ObjectIDFromHex(transactionId)
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

		err = h.store.UpdateTransactionById(existingTransaction)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating transaction %v", err))
			return
		}



		utils.WriteJSON(w, http.StatusOK, existingTransaction)
		
}

// @Summary Delete Transaction by Id
// @tags Transaction
// @Param id path string true "Transaction ID"
// @Success 200 {boolean} bool
// @Router /transaction/{transactionId} [DELETE]
func (h *Handler) HandleDeleteTransactionById(w http.ResponseWriter, r *http.Request) {
		transactionId := r.PathValue("transactionId")
		if transactionId == "" {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("provide non empty transactionid"))
			return
		}

		objectId, err := primitive.ObjectIDFromHex(transactionId)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid transactionid %v", err))
			return
		}

		err = h.store.DeleteTransactionById(objectId)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error deleting transaction %v", err))
			return
		}
	
		utils.WriteJSON(w, http.StatusOK, true)

}

// @Summary Gets Transactions by Query
// @tags Transaction
// @Produce json
// @Param   userId     query    string      false "Filter transactions by userId"
// @Success 200 {array} types.Transaction
// @Router /transaction [GET]
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


// @Summary Create Transaction 
// @tags Transaction
// @Produce json
// @Accept  json
// @Param   transaction  body     types.CreateTransactionPayload  true  "Transaction to Create"
// @Success 200 {object} types.Transaction
// @Router /transaction [POST]
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
	user, err := h.userStore.GetUserById(payload.UserId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user does not exist %v", err))
		return
	}

	
	newTransaction := types.Transaction{
		UserId: 					user.Id,	
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

	newTransaction.Id = id
	utils.WriteJSON(w, http.StatusOK, newTransaction)
}