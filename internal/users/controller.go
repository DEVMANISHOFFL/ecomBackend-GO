package users

import (
	"database/sql"
	"ecom/internal/cart"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ProfileWithCart struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Role      string          `json:"role"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Cart      []cart.CartItem `json:"cart"` 
}

func CreateUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := CreateUserService(db, u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func GetUsersController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := GetAllUsersService(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}
}

func GetUserByIdController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		user, err := GetUserByIdService(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func DeleteUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]

		user, err := FetchUserById(db, id)
		if err != nil {
			http.Error(w, `{"error": "Database error while fetching user"}`, http.StatusInternalServerError)
			return
		}

		if user == nil {
			http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
			return
		}

		deleted, err := DeleteUserService(db, id)
		if err != nil {
			http.Error(w, `{"error": "Error deleting user"}`, http.StatusInternalServerError)
			return
		}

		if !deleted {
			http.Error(w, `{"error": "User could not be deleted"}`, http.StatusNotFound)
			return
		}

		// Success response
		json.NewEncoder(w).Encode(map[string]any{
			"message": "User deleted successfully",
			"user":    user,
		})
	}
}

func UpdateUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedUser, err := UpdateUserService(db, id, u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"message": "User updated", "user": updatedUser})
	}
}
