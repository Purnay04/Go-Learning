package repo

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	AddUser(*User) (*User, error)
	GetUserById(int) (*User, error)
	AuthenticateUser(string, string) (*User, error)
	GetUserByUsernameOrEmail(string, string) (*User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (userRepo *UserRepositoryImpl) AuthenticateUser(email string, password string) (*User, error) {
	query := `SELECT user_id, name, email FROM user where email = $1 and password = $2`
	interpolateVals := []interface{}{
		email,
		password,
	}
	result, err := userRepo.db.Query(query, interpolateVals...)

	if err != nil {
		return nil, err
	}

	users := []User{}
	for result.Next() {
		var user User
		if err := result.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("wrong credentials")
	} else if len(users) > 1 {
		return nil, fmt.Errorf("duplicate users are present")
	}
	return &users[0], nil
}

func (userRepo *UserRepositoryImpl) AddUser(newUser *User) (*User, error) {
	user, err := userRepo.GetUserByUsernameOrEmail(newUser.Name, newUser.Email)
	if err != nil {
		return nil, err
	} else if user != nil {
		return nil, fmt.Errorf("User already exist with username or email")
	}

	query := `INSERT INTO user(name, email, password) VALUES ($1, $2, $3) RETURING user_id`
	interpolateVals := []interface{}{newUser.Name, newUser.Email, newUser.Password}

	row := userRepo.db.QueryRow(query, interpolateVals...)
	err = row.Scan(&newUser.Id)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (userRepo *UserRepositoryImpl) GetUserById(id int) (*User, error) {
	user := &User{}
	query := `SELECT name, email FROM user WHERE user_id = $1`
	err := userRepo.db.QueryRow(query, id).Scan(&user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepo *UserRepositoryImpl) GetUserByUsernameOrEmail(username string, email string) (*User, error) {
	user := &User{}
	query := `SELECT user_id, name, email FROM user WHERE name = $1 or email = $2`
	err := userRepo.db.QueryRow(query, username, email).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
