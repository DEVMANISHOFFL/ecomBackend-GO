package products

import (
	"database/sql"
	"ecom/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProductController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u Product
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body"))
			return
		}
		resp, err := CreateProductService(db, u)
		if err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}
}

func GetProductsController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := GetAllProductsService(db)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func GetProductByIdController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		product, err := GetProductByIdService(db, id)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		if product == nil {
			utils.SendJSONError(w, http.StatusFound, fmt.Errorf("user not found"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func DeleteProductController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		product, err := FetchProductById(db, id)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		if product == nil {
			utils.SendJSONError(w, http.StatusNotFound, fmt.Errorf("product not found"))
			return
		}
		deleted, err := DeleteProductService(db, id)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		if !deleted {
			utils.SendJSONError(w, http.StatusNotFound, fmt.Errorf("product not found"))
			return
		}
		json.NewEncoder(w).Encode(map[string]any{"message": "Product deleted successfully", "product": product})
	}
}

func UpdateProductController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var p Product
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			utils.SendJSONError(w, http.StatusBadRequest, fmt.Errorf("invalid JSON body"))
			return
		}
		updatedProduct, err := UpdateProductService(db, id, p)
		if err != nil {
			utils.SendJSONError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]any{"message": "Product updated successfully", "Product": updatedProduct})
	}
}
