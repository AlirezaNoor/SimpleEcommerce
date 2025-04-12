package repository

import (
    "ecommerce/internal/models"
)

func CreateOrder(order *models.Order) error {
    // درج سفارش و دریافت شناسه سفارش
    err := DB.QueryRow(
        "INSERT INTO orders (total_price) VALUES ($1) RETURNING id",
        order.TotalPrice,
    ).Scan(&order.ID)
    if err != nil {
        return err
    }

    // درج هر یک از محصولات سفارش داده شده در جدول order_products
    for _, productID := range order.ProductIDs {
        if _, err := DB.Exec(
            "INSERT INTO order_products (order_id, product_id) VALUES ($1, $2)",
            order.ID, productID,
        ); err != nil {
            // در اینجا می‌توانید منطق rollback یا مدیریت خطا برای هر محصول را اضافه کنید
            return err
        }
    }
    return nil
}
