package seeder

import "github.com/tavvfiq/cafe-rest-api-gorm/database/model"

// CreateLevel create level seeders
func CreateLevel() *[]model.Level {
	_levels := []model.Level{{
		ID:   1,
		Name: "Customer",
	}, {
		ID:   2,
		Name: "Cashier",
	}, {
		ID:   3,
		Name: "Supervisor",
	}, {
		ID:   4,
		Name: "Admin",
	}, {ID: 5,
		Name: "Super Admin"},
	}
	return &_levels
}
