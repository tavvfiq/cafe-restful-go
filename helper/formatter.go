package helper

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tavvfiq/cafe-rest-api-gorm/database/model"
)

// ToAuthFormat format user model to auth response
func ToAuthFormat(user *model.User) interface{} {
	type response struct {
		ID        uuid.UUID `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		LevelID   uint8     `json:"level_id"`
	}

	return response{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		LevelID:   user.LevelID,
	}
}

// FromQuery extract query params data
func FromQuery(ctx echo.Context, qm map[string]string) {
	menuQuery := []string{"search", "filter", "sortby", "order", "page", "limit"}
	for _, key := range menuQuery {
		qm[key] = ctx.QueryParam(key)
	}
}
