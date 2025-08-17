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
