package users

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(db *sql.DB, u User) (UserResponse, error) {
	if u.Name == "" || u.Email == "" || u.Password == "" {
		return UserResponse{}, errors.New("name, email and password are required")
	}
	if u.Role == "" {
		u.Role = RoleCustomer
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, err
	}
	u.Password = string(hashed)

	id, err := InsertUser(db, u)
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:    id,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}, nil
}

func GetAllUsersService(db *sql.DB) ([]UserResponse, error) {
	return FetchUsers(db)
}

func GetUserByIdService(db *sql.DB, id string) (*UserResponse, error) {
	return FetchUserById(db, id)
}

func DeleteUserService(db *sql.DB, id string) (bool, error) {
	return DeleteUser(db, id)
}

func UpdateUserService(db *sql.DB, id string, u User) (*UserResponse, error) {
	return UpdateUser(db, id, u)
}
