package user

import (
	"fmt"
	"net/http"

	"github.com/bjhermid/go-api-1/services/auth"
	"github.com/bjhermid/go-api-1/types"
	"github.com/bjhermid/go-api-1/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validated the payload
	//add to the types/model after `json` format
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	user, err := h.store.GetUserByEmail(payload.Email)
	if err == nil && user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found, invalid email or password"))
		return
	}

	if !auth.ComparePassword(user.Password,[]byte(payload.Password)){
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user not found, invalid email or password"))
		return
	}
	utils.WriteJson(w,http.StatusOK, map[string]string{"token":""})
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validated the payload
	//add to the types/model after `json` format
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//check if user exist
	user, err := h.store.GetUserByEmail(payload.Email)
	if err == nil && user != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exist", payload.Email))
		return
	}

	//hash password
	hashPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	//if not create a new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusCreated, nil)
}
