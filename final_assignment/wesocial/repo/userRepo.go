package repo

import "database/sql"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	AddUser(*User) (*User, error)
	GetUserById(int) (*User, error)
	AuthenticateUser(string, string) (bool, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (userRepo *UserRepositoryImpl) AuthenticateUser(email string, password string) (bool, error) {
	query := `SELECT COUNT(user_id) FROM user where email = $1 and password = $2`
	interpolateVals := []interface{}{
		email,
		password,
	}
	var count int
	err := userRepo.db.QueryRow(query, interpolateVals...).Scan(&count)

	if err != nil {
		return false, err
	}

	if count == 1 {
		return true, nil
	}
	return false, nil
}

func (userRepo *UserRepositoryImpl) AddUser(newUser *User) (*User, error) {
	query := `INSERT INTO user(name, email, password) VALUES ($1, $2, $3) RETURING user_id`
	interpolateVals := []interface{}{newUser.Name, newUser.Email, newUser.Password}

	row := userRepo.db.QueryRow(query, interpolateVals...)
	err := row.Scan(&newUser.Id)

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
