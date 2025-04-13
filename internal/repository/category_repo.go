package repository

import "ecommerce/internal/models"

// CreateCategory ایجاد دسته‌بندی جدید
func CreateCategory(category *models.Category) error {
	err := DB.QueryRow(
		"INSERT INTO categories (name) VALUES ($1) RETURNING id",
		category.Name,
	).Scan(&category.ID)
	return err
}

// GetCategories دریافت لیست دسته‌بندی‌ها
func GetCategories() ([]models.Category, error) {
	rows, err := DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
			return nil, err
		}
		categories = append(categories, cat)
	}
	return categories, nil
}
