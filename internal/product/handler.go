package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products/all", h.GetAll).Methods("GET")
	r.HandleFunc("/products/{id}", h.GetByID).Methods("GET")
	r.HandleFunc("/products/create", h.Create).Methods("POST")
	r.HandleFunc("/products/update/{id}", h.Update).Methods("PUT")
	r.HandleFunc("/products/delete/{id}", h.Delete).Methods("DELETE")
}

// GetAllProducts godoc
// @Summary دریافت همه محصولات
// @Description این متد همه محصولات را برمی‌گرداند
// @Tags محصولات
// @Produce  json
// @Success 200 {array} product.Product
// @Router /products/all [get]
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

// GetProductByID godoc
// @Summary دریافت محصول با شناسه
// @Description این متد یک محصول را با استفاده از ID برمی‌گرداند
// @Tags محصولات
// @Produce  json
// @Param id path int true "شناسه محصول"
// @Success 200 {object} product.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)
	product, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// CreateProduct godoc
// @Summary ایجاد محصول جدید
// @Description این متد یک محصول جدید ایجاد می‌کند
// @Tags محصولات
// @Accept  json
// @Produce  json
// @Param product body product.Product true "اطلاعات محصول"
// @Success 201 {object} product.Product
// @Failure 400 {object} map[string]string
// @Router /products/create [post]
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

// UpdateProduct godoc
// @Summary بروزرسانی محصول
// @Description این متد یک محصول را بروزرسانی می‌کند
// @Tags محصولات
// @Accept  json
// @Produce  json
// @Param id path int true "شناسه محصول"
// @Param product body product.Product true "اطلاعات محصول"
// @Success 200 {object} product.Product
// @Failure 400 {object} map[string]string
// @Router /products/update/{id} [put]

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var p Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	p.ID = id
	if err := h.service.Update(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(p)
}

// DeleteProduct godoc
// @Summary حذف محصول
// @Description این متد یک محصول را حذف می‌کند
// @Tags محصولات
// @Param id path int true "شناسه محصول"
// @Success 204 "حذف موفقیت‌آمیز"
// @Failure 500 {object} map[string]string
// @Router /products/delete/{id} [delete]
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if err := h.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
