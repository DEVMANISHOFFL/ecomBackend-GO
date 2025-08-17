package users

import "database/sql"

func InsertUser(db *sql.DB, u User) (string, error) {
	var id string
	err := db.QueryRow(`
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, COALESCE(NULLIF($4, ''), 'customer'), NOW(), NOW())
		RETURNING id`,
		u.Name, u.Email, u.Password, u.Role,
	).Scan(&id)

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
