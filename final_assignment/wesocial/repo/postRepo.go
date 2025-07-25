package repo

import (
	"database/sql"
	"time"
)

type Post struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedBy int       `json:"created_by"`
	CreatedOn time.Time `json:"created_on"`
}

type PostRepository interface {
	CreatePost(*Post) (*Post, error)
	GetAllPost() ([]Post, error)
	GetPostByUser(int) ([]Post, error)
}

type PostRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &PostRepositoryImpl{db: db}
}

func (postRepo *PostRepositoryImpl) CreatePost(newPost *Post) (*Post, error) {
	query := `INSERT INTO post (title, content, created_by, created_on) values ($1, $2, $3, $4) RETURNING post_id`
	interpolateVals := []interface{}{
		newPost.Title,
		newPost.Content,
		newPost.CreatedBy,
		newPost.CreatedOn,
	}
	err := postRepo.db.QueryRow(query, interpolateVals...).Scan(&newPost.Id)

	if err != nil {
		return nil, err
	}
	return newPost, nil
}

func (postRepo *PostRepositoryImpl) GetAllPost() ([]Post, error) {
	query := `SELECT title, content, created_by, created_on FROM post`
	result, err := postRepo.db.Query(query)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for result.Next() {
		var post Post
		if err := result.Scan(&post.Title, &post.Content, &post.CreatedBy, &post.CreatedOn); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (postRepo *PostRepositoryImpl) GetPostByUser(userId int) ([]Post, error) {
	query := `SELECT title, content, created_by, created_on FROM post WHERE created_by = $1`
	result, err := postRepo.db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	posts := []Post{}
	for result.Next() {
		var post Post
		if err := result.Scan(&post.Title, &post.Content, &post.CreatedBy, &post.CreatedOn); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
