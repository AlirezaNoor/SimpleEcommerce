package handlers

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/models"
	"ecommerce/internal/repository"
)

// GetProducts godoc
// @Summary دریافت لیست محصولات
// @Description لیست تمامی محصولات موجود در فروشگاه را برمی‌گرداند.
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {string} string "Internal Server Error"
// @Router /products [get]
func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := repository.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// CreateProduct godoc
// @Summary افزودن محصول جدید
// @Description ایجاد و ثبت یک محصول جدید در فروشگاه.
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product Data"
// @Success 200 {object} models.Product
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products [post]
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := repository.CreateProduct(&product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// CreateOrder godoc
// @Summary ثبت سفارش جدید
// @Description ایجاد یک سفارش جدید با لیست محصولات و مبلغ کل سفارش.
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order Data"
// @Success 200 {object} models.Order
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /orders [post]
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := repository.CreateOrder(&order); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
