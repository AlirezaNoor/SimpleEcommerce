package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "ecommerce/docs" // import docs for swagger

	"ecommerce/internal/product"
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

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repository.InitDB(dsn)

	router := mux.NewRouter()

	// مسیر Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// روت‌های محصولات
	repo := product.NewRepository(db)
	service := product.NewService(repo)
	productHandler := product.NewHandler(service)
	productHandler.RegisterRoutes(router)

	log.Println("سرور روی پورت 8080 در حال اجراست...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
