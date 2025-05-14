package transaction

import (
	"fmt"
	"net/http"
	"psp-dashboard-be/types"
	"psp-dashboard-be/utils"
	"time"

	"github.com/go-playground/validator/v10"
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
}


func (h *Handler) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {

		utils.WriteJSON(w, http.StatusOK, "Get Transactions")
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