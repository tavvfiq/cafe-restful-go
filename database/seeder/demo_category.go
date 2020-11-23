package seeder

import "github.com/tavvfiq/cafe-rest-api-gorm/database/model"

// CreateCategory create level seeders
func CreateCategory() *[]model.Category {
	_categories := []model.Category{{
		ID:   1,
		Name: "Main Course",
	}, {
		ID:   2,
		Name: "Dessert",
	}, {
		ID:   3,
		Name: "Beverage",
	}, {
		ID:   4,
		Name: "Snack",
	}}
	return &_categories
}
