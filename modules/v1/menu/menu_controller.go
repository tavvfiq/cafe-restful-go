package menu

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
	"github.com/tavvfiq/cafe-rest-api-gorm/database"
	"github.com/tavvfiq/cafe-rest-api-gorm/helper"
)

func addMenu(ctx echo.Context) error {
	qm := make(map[string]string)
	helper.FromQuery(ctx, qm)

	page, err := strconv.ParseInt(qm["page"], 10, 32)

	if err != nil {
		log.Fatal(errors.Wrap(err, "convert page to int error"))
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "internal server error"})
	}

	perPage, err := strconv.ParseInt(qm["limit"], 10, 32)

	if err != nil {
		log.Fatal(errors.Wrap(err, "convert limit to int error"))
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "internal server error"})
	}

	database.Db.Where("name LIKE ?", fmt.Sprintf("%s%"))

	return nil
}
