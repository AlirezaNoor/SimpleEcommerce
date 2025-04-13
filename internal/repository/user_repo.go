package repository

import (
	"ecommerce/internal/models"
	"time"
)

// CreateUser ایجاد کاربر جدید
func CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	return DB.QueryRow(query, user.Name, user.Email, user.Password, time.Now()).Scan(&user.ID)
}

// GetUsers دریافت لیست کاربران
func GetUsers() ([]models.User, error) {
	rows, err := DB.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
