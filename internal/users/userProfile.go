package users

import (
	"database/sql"
	"ecom/internal/cart"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProfileController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := GetUserByIdService(db, id)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		userCart, err := cart.GetCartByUserIDService(db, id)
		if err != nil {
			http.Error(w, "Failed to fetch cart: "+err.Error(), http.StatusInternalServerError)
			return
		}

		profile := ProfileWithCart{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Cart:      userCart.Items,
		}

		json.NewEncoder(w).Encode(profile)
	}
}
