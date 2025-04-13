package models

import "time"

// User مدل کاربر
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // در خروجی JSON نمایش داده نمی‌شود
	CreatedAt time.Time `json:"created_at"`
}
