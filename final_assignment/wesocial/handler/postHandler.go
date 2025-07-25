package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wesocial/repo"
	"wesocial/service"
)

type PostHandler struct {
	postService service.PostService
}

func NewPostHandler(postService service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

func (h *PostHandler) AddNewPost(w http.ResponseWriter, r *http.Request) {
	post := &repo.Post{}

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Request body is not valid or in incorrect format", http.StatusBadRequest)
		return
	}

	post, err = h.postService.AddPost(post)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// byUserIdStr := vars["user_id"]
	byUserIdStr := r.URL.Query().Get("user_id")

	byUserId, err := strconv.Atoi(byUserIdStr)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	posts, err := h.postService.GetAllPost(byUserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
