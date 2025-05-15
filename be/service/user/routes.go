package user

import (
	"fmt"
	"net/http"
	"psp-dashboard-be/types"
	"psp-dashboard-be/utils"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	router.HandleFunc("DELETE /user/{userId}", h.HandleDeleteUserById)
	router.HandleFunc("PUT /user/{userId}", h.HandleUpdateUserById)	
}

// @Summary Updates User by Id
// @tags User
// @Produce json
// @Accept  json
// @Param   user  body     types.UpdateUserPayload  true  "User to update"
// @Param id path string true "User ID"
// @Success 200 {object} types.User
// @Router /user/{userId} [PUT]
func (h *Handler) HandleUpdateUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userId")
	if userId == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("provide non empty userId"))
		return
	}

	
	var payload types.UpdateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	existingUser, err := h.store.GetUserById(userId)
	if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error fetching user %v", err))
			return
	}

	existingUser.Email = payload.Email
	existingUser.Name = payload.Name
	existingUser.Notes = payload.Notes

	err = h.store.UpdateUserById(*existingUser)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error updating user %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, existingUser)
}


// @Summary Delete User by Id
// @tags User
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {boolean} bool
// @Router /user/{userId} [DELETE]
func (h *Handler) HandleDeleteUserById(w http.ResponseWriter, r *http.Request) {
 	userId := r.PathValue("userId")
	if userId == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("provide non empty userId"))
		return
	}

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid userId %v", err))
		return
	}

	err = h.store.DeleteUserById(objectId)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error deleting user %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, true)

}

// @Summary Get all Users
// @tags User
// @Produce json
// @Success 200 {array} types.User
// @Router /user [GET]
func (h *Handler) HandleGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetUsersWithTransactions()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("error getting users %v", err))
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

// @Summary Creates a new User
// @tags User
// @Produce json
// @Accept  json
// @Param   user  body     types.UpdateUserPayload  true  "User to update"
// @Success 200 {object} types.User
// @Router /user [POST]
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