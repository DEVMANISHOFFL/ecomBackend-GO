package users

import (
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid Json body", http.StatusBadRequest)
			return
		}
		if u.Name == "" || u.Email == "" || u.Password == "" {
			http.Error(w, "Name,Email and Password are required", http.StatusBadRequest)
			return
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error securing password", http.StatusInternalServerError)
			return
		}

		err = db.QueryRow(`
    INSERT INTO users (name, email, password, role, created_at, updated_at)
    VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`,
			u.Name, u.Email, hashed, u.Role,
		).Scan(&u.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp := UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.Role,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}
