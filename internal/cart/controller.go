package cart

import (
	"database/sql"
	"ecom/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateCartController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u Cart
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body"))
			return
		}
		res, err := CreateCartService(db, u)
		if err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(res)
	}
}

func GetCartByIdController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		fmt.Println(id)
		cart, err := FetchCartById(db, id)
		fmt.Println(cart)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		if cart == nil {
			utils.SendJSONError(w, http.StatusNotFound, fmt.Errorf("cart not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cart)
	}
}
