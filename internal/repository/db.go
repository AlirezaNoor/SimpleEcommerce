package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB اتصال به پایگاه داده را برقرار کرده و جداول را ایجاد می‌کند
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("خطا در اتصال به پایگاه داده: ", err)
	}

	// بررسی اتصال
	if err = DB.Ping(); err != nil {
		log.Fatal("خطا در پینگ پایگاه داده: ", err)
	}

	createTables()
}

func createTables() {
	// جدول دسته‌بندی‌ها
	categoriesTable := `
	CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	);`
	if _, err := DB.Exec(categoriesTable); err != nil {
		log.Fatal("خطا در ایجاد جدول categories: ", err)
	}

	// جدول محصولات با CategoryID
	productsTable := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL,
		category_id INTEGER REFERENCES categories(id)
	);`
	if _, err := DB.Exec(productsTable); err != nil {
		log.Fatal("خطا در ایجاد جدول products: ", err)
	}

	// جدول سفارشات
	ordersTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		total_price NUMERIC(10,2) NOT NULL
	);`
	if _, err := DB.Exec(ordersTable); err != nil {
		log.Fatal("خطا در ایجاد جدول orders: ", err)
	}

	// جدول ارتباط سفارشات و محصولات
	orderProductsTable := `
	CREATE TABLE IF NOT EXISTS order_products (
		order_id INTEGER REFERENCES orders(id),
		product_id INTEGER REFERENCES products(id),
		PRIMARY KEY (order_id, product_id)
	);`
	if _, err := DB.Exec(orderProductsTable); err != nil {
		log.Fatal("خطا در ایجاد جدول order_products: ", err)
	}

	// جدول کاربران
	usersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	if _, err := DB.Exec(usersTable); err != nil {
		log.Fatal("خطا در ایجاد جدول users: ", err)
	}
}
