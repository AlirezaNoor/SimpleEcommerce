package repository

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB اتصال به پایگاه داده را برقرار کرده و در صورت نیاز جداول را ایجاد می‌کند
func InitDB(dataSourceName string) {
    var err error
    DB, err = sql.Open("postgres", dataSourceName)
    if err != nil {
        log.Fatal("خطا در اتصال به پایگاه داده: ", err)
    }

    // بررسی اتصال به پایگاه داده
    if err = DB.Ping(); err != nil {
        log.Fatal("خطا در پینگ پایگاه داده: ", err)
    }

    // ایجاد جداول در صورت عدم وجود
    createTables()
}

func createTables() {
    // ایجاد جدول محصولات
    productsTable := `
    CREATE TABLE IF NOT EXISTS products (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        price NUMERIC(10,2) NOT NULL
    );`

    if _, err := DB.Exec(productsTable); err != nil {
        log.Fatal("خطا در ایجاد جدول products: ", err)
    }

    // ایجاد جدول سفارشات
    ordersTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id SERIAL PRIMARY KEY,
        total_price NUMERIC(10,2) NOT NULL
    );`

    if _, err := DB.Exec(ordersTable); err != nil {
        log.Fatal("خطا در ایجاد جدول orders: ", err)
    }

    // ایجاد جدول میانجی برای ارتباط سفارشات و محصولات
    orderProductsTable := `
    CREATE TABLE IF NOT EXISTS order_products (
        order_id INTEGER REFERENCES orders(id),
        product_id INTEGER REFERENCES products(id),
        PRIMARY KEY (order_id, product_id)
    );`

    if _, err := DB.Exec(orderProductsTable); err != nil {
        log.Fatal("خطا در ایجاد جدول order_products: ", err)
    }
}
