package handler

import (
	"encoding/json"
	"net/http"
	"wesocial/repo"
	"wesocial/service"
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

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

}

func NewUserHandler
