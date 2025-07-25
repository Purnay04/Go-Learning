package main

import (
	"database/sql"
	"net/http"
	"strings"
	"wesocial/handler"
	"wesocial/repo"
	"wesocial/service"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var tokenPrefix = "Bearer "

func tokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, tokenPrefix) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		token := strings.TrimPrefix(authorizationHeader, tokenPrefix)
		tokenStruct, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return token, nil
		})
		if err != nil || !tokenStruct.Valid {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func getUserRouteHandler(db *sql.DB) *handler.UserHandler {
	repo := repo.NewUserRepository(db)
	userService := service.NewUserService(repo)
	return handler.NewUserHandler(userService)
}

func getPostRouteHandler(db *sql.DB) *handler.PostHandler {
	postRepo := repo.NewPostRepository(db)
	userRepo := repo.NewUserRepository(db)
	postService := service.NewPostService(postRepo, userRepo)
	return handler.NewPostHandler(postService)
}

func main() {
	connectionStr := "user=postgres password=Cloud@123$ dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userHandler, postHandler := getUserRouteHandler(db), getPostRouteHandler(db)

	router := mux.NewRouter()

	// User routes
	router.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	// Post routes
	router.HandleFunc("/posts", tokenMiddleware(postHandler.AddNewPost)).Methods("POST")
	router.HandleFunc("/posts", tokenMiddleware(postHandler.GetPosts)).Methods("GET")

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}
