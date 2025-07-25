package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"wesocial/repo"
	"wesocial/service"
	"wesocial/utils"
)

type UserHandler struct {
	userService service.UserService
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user repo.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	created, err := h.userService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(created)
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	cred := struct {
		Email    string
		Password string
	}{}

	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		http.Error(w, "please provide the proper credentials details", http.StatusBadRequest)
		return
	}

	user, err := h.userService.LoginUser(cred.Email, cred.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := utils.CreateToken(map[string]any{
		"username": user.Name,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tokenResponse{
		Token: token,
	})
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
