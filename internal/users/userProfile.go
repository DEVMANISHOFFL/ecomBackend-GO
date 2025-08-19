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
			Role:      string(user.Role), // convert Role type to string
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Cart:      userCart.Items,
		}

		json.NewEncoder(w).Encode(profile)
	}
}

func GetProfileCartController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		if id == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		cart, err := cart.GetCartByUserIDService(db, id)
		if err != nil {
			http.Error(w, "failed to fetch cart: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if len(cart.Items) == 0 {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]any{
				"user_id": id,
				"cart":    []any{},
				"message": "Cart is empty",
			})
			return
		}
		json.NewEncoder(w).Encode(cart)
	}

}
