package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"ecom/pkg/utils"

	"github.com/gorilla/mux"
)

func CreateUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body"))
			return
		}

		resp, err := CreateUserService(db, u)
		if err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, err)
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
			utils.SendJSONError(w, http.StatusInternalServerError, err)
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
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		if user == nil {
			utils.SendJSONError(w, http.StatusNotFound, fmt.Errorf("user not found"))
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
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}

		if user == nil {
			utils.SendJSONError(w, http.StatusNotFound, fmt.Errorf("User not found"))
			return
		}

		deleted, err := DeleteUserService(db, id)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		if !deleted {
			utils.SendJSONError(w, http.StatusNotFound, fmt.Errorf("User not found"))
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"message": "User Deleted Successfully", "user": user})
	}
}

func UpdateUserController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body"))
			return
		}
		UpdatedUser, err := UpdateUserService(db, id, u)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"message": "User updated", "user": UpdatedUser})
	}
}
