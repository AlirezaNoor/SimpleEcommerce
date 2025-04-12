package repository

import (
    "ecommerce/internal/models"
)

// دریافت لیست محصولات از پایگاه داده
func GetProducts() ([]models.Product, error) {
    rows, err := DB.Query("SELECT id, name, price FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var p models.Product
        if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
            return nil, err
        }
        products = append(products, p)
    }
    return products, nil
}

// افزودن محصول جدید به پایگاه داده
func CreateProduct(product *models.Product) error {
    err := DB.QueryRow(
        "INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id",
        product.Name, product.Price,
    ).Scan(&product.ID)
    return err
}
