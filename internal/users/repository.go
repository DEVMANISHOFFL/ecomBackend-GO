package users

import "database/sql"

func InsertUser(db *sql.DB, u User) (int, error) {
	var id int
	err := db.QueryRow(`
		INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`,
		u.Name, u.Email, u.Password, u.Role,
	).Scan(&id)
	return id, err
}

func FetchUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id,name,email,role,created_at,updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
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
