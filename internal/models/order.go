package models

// مدل سفارش
type Order struct {
    ID         int     `json:"id"`
    ProductIDs []int   `json:"product_ids"`
    TotalPrice float64 `json:"total_price"`
}
