package users

import (
	"database/sql"
	"fmt"
)

func InsertUser(db *sql.DB, u User) (string, error) {
	var id string
	err := db.QueryRow(`
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, COALESCE(NULLIF($4, ''), 'customer'), NOW(), NOW())
		RETURNING id`,
		u.Name, u.Email, u.Password, u.Role,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("failed to insert user: %w", err)
	}

	return id, err
}

func FetchUsers(db *sql.DB) ([]UserResponse, error) {
	rows, err := db.Query("SELECT id,name,email,role,created_at,updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserResponse
	for rows.Next() {
		var u UserResponse
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func FetchUserById(db *sql.DB, id string) (*UserResponse, error) {
	var u UserResponse
	err := db.QueryRow("SELECT id, name, email, role, created_at, updated_at FROM users WHERE id=$1", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func DeleteUserById(db *sql.DB, id string) (bool, error) {
	res, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return false, fmt.Errorf("failed to delete user: %w", err)
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func UpdateUser(db *sql.DB, id string, u User) (*UserResponse, error) {
	_, err := db.Exec(`
		UPDATE users
		SET name=$1, email=$2, password=$3, role=$4,updated_at=NOW()
		WHERE id=$5`, u.Name, u.Email, u.Password, u.Role, id)

	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return FetchUserById(db, id)
}
