package auth

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tavvfiq/cafe-rest-api-gorm/apiutils"
	"github.com/tavvfiq/cafe-rest-api-gorm/database"
	"github.com/tavvfiq/cafe-rest-api-gorm/database/model"
	"github.com/tavvfiq/cafe-rest-api-gorm/helper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	salt = 10
)

// RegisterHandler handler register
func RegisterHandler(ctx echo.Context) error {
	r := new(model.User)

	// bind http body to model struct
	if err := ctx.Bind(r); err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "Error registering user"})
	}

	// hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), salt)
	if err != nil {
		panic(err)
	}

	// generate random ID
	r.ID = uuid.New()

	// create user
	if err := database.Db.Model(&model.User{}).Create(&model.User{ID: r.ID, FirstName: r.FirstName, LastName: r.LastName, Email: r.Email, Password: string(hashedPassword), LevelID: r.LevelID}).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "user already exist"})
	}

	//sign token
	payload := jwt.MapClaims{
		"id":         r.ID,
		"level_id":   r.LevelID,
		"created_at": time.Now()}
	tokenString, err := apiutils.GenerateJwt([]byte(os.Getenv("jwt_secret")), payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "Generate token error"})
	}

	return ctx.JSON(http.StatusOK, helper.AuthResponse{Status: http.StatusOK, Message: "register success", Data: helper.ToAuthFormat(r), Token: tokenString})
}

// LoginHandler handle login
func LoginHandler(ctx echo.Context) error {
	l := new(model.User)

	// bind http body to model struct
	if err := ctx.Bind(l); err != nil {
		log.Fatal(err)
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "error log in user"})
	}
	// var to store database result
	result := model.User{}
	// check user from database
	if err := database.Db.Model(&model.User{}).Where("email = ?", l.Email).First(&result).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.JSON(http.StatusNoContent, helper.ErrorResponse{Status: http.StatusNoContent, Error: "user not found"})
	} else if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "internal server error"})
	}

	// compare hashed password with login password
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(l.Password)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "wrong password"})
	}

	//sign token
	payload := jwt.MapClaims{
		"id":         result.ID,
		"level_id":   result.LevelID,
		"created_at": time.Now()}
	tokenString, err := apiutils.GenerateJwt([]byte(os.Getenv("jwt_secret")), payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse{Status: http.StatusInternalServerError, Error: "Generate token error"})
	}

	return ctx.JSON(http.StatusOK, helper.AuthResponse{Status: http.StatusOK, Message: "login success", Data: helper.ToAuthFormat(&result), Token: tokenString})
}
