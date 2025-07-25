package service

import (
	"fmt"
	"wesocial/repo"
)

type PostService interface {
	AddPost(*repo.Post) (*repo.Post, error)
	GetAllPost(int) ([]repo.Post, error)
}

type PostServiceImpl struct {
	postRepo repo.PostRepository
	userRepo repo.UserRepository
}

func NewPostService(postRepo repo.PostRepository, userRepo repo.UserRepository) PostService {
	return &PostServiceImpl{
		postRepo: postRepo,
		userRepo: userRepo,
	}
}

func (postService *PostServiceImpl) AddPost(newPost *repo.Post) (*repo.Post, error) {
	if newPost.Title == "" {
		return nil, fmt.Errorf("post title is missing")
	}
	if newPost.Content == "" {
		return nil, fmt.Errorf("post contents are missing")
	}
	if newPost.CreatedBy != 0 {
		return nil, fmt.Errorf("post creator id should be present")
	}
	if newPost.CreatedOn.IsZero() {
		return nil, fmt.Errorf("post date should be present")
	}

	return postService.postRepo.CreatePost(newPost)
}

func (postService *PostServiceImpl) GetAllPost(byUserID int) ([]repo.Post, error) {
	if byUserID != 0 {
		return postService.postRepo.GetPostByUser(byUserID)
	} else {
		return postService.postRepo.GetAllPost()
	}
}
