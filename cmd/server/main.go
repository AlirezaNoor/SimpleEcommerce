package main

import (
	"log"
	"net/http"
	"os"

	_ "ecommerce/docs" // import docs for swagger
	"ecommerce/internal/handlers"
	"ecommerce/internal/repository"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Ecommerce API
// @version 1.0
// @description API for a simple ecommerce platform including users and categories.
// @host localhost:8080
// @BasePath /

func main() {
	// خواندن DSN از متغیر محیطی یا استفاده از مقدار پیش‌فرض
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:55375447@localhost:5432/ecommerce?sslmode=disable"
	}
	repository.InitDB(dsn)

	router := mux.NewRouter()

	// مسیر Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// روت‌های محصولات
	router.HandleFunc("/products", handlers.GetProducts).Methods("GET")
	router.HandleFunc("/products", handlers.CreateProduct).Methods("POST")
	// روت سفارشات
	router.HandleFunc("/orders", handlers.CreateOrder).Methods("POST")
	// روت‌های کاربران
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	// (می‌توانید روت دریافت کاربران را نیز اضافه کنید)
	// روت‌های دسته‌بندی‌ها (در صورت نیاز)
	// router.HandleFunc("/categories", handlers.CreateCategory).Methods("POST")

	log.Println("سرور روی پورت 8080 در حال اجراست...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
