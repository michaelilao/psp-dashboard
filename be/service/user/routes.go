package user

import (
	"fmt"
	"net/http"
	"psp-dashboard-be/types"
	"psp-dashboard-be/utils"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	store types.UserStore
}


func NewHandler(store types.UserStore) *Handler{
	return &Handler{
		store: store,
	}
}


func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /user", h.HandleGetUsers)
	router.HandleFunc("POST /user", h.HandleCreateUser)
}

func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := h.store.GetUsers()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting users %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err := h.store.GetUserByEmail(payload.Email)
	if err != mongo.ErrNoDocuments {
		utils.WriteError(w, http.StatusConflict, fmt.Errorf("user with that email already exists"))
		return
	}

	newUser := types.User{
		Name: payload.Name,
		Email: payload.Email,
		Notes: payload.Notes,
	}

	id, err := h.store.InsertUser(newUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error inserting new user %v", err))
		return
	}

	newUser.Id = id
	utils.WriteJSON(w, http.StatusOK, newUser)
}