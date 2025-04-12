package main

import (
	"log"
	"net/http"
	"os"

	"ecommerce/internal/handlers"
	"ecommerce/internal/repository"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "ecommerce/docs"
)

// @title Ecommerce API
// @version 1.0
// @description API for a simple ecommerce platform.
// @host localhost:8080
// @BasePath /

func main() {
	// خواندن DSN از متغیر محیطی یا استفاده از مقدار پیش‌فرض
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:55375447@localhost:5432/ecommerce?sslmode=disable"
	}

	// مقداردهی اولیه به پایگاه داده
	repository.InitDB(dsn)

	router := mux.NewRouter()

	// مسیر Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// روت‌های محصولات
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")

	// روت ثبت سفارش
	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")

	log.Println("سرور روی پورت 8080 در حال اجراست...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
