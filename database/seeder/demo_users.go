package seeder

import (
	"github.com/google/uuid"
	"github.com/tavvfiq/cafe-rest-api-gorm/database/model"
)

// CreateUser user seeders
func CreateUser() *[]model.User {
	_users := []model.User{{
		ID:          uuid.New(),
		FirstName:   "Taufiq Widi",
		LastName:    "Nugroho",
		PhoneNumber: "081284544654",
		Email:       "taufiqwidinugroho@gmail.com",
		Password:    "$2b$10$0XwKUYDcF.RywZPNhEmYSuZKncL7pCeYSNwUOihk/SZ4FF50hPVKu",
		LevelID:     4,
	}, {
		ID:        uuid.New(),
		FirstName: "Razor",
		LastName:  "Back",
		Email:     "razorback@mail.com",
		Password:  "$2b$10$goPXXEo5DR9YkLAEeoItnu6WLw1ZWIvk./nt6.lYDwLBTpfJ/qRTy",
		LevelID:   1,
	}}

	return &_users
}
