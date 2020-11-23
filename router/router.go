package router

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tavvfiq/cafe-rest-api-gorm/modules/v1/auth"
)

// Start the router
func Start() {
	e := echo.New()
	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}\turi=${uri}\t\tstatus=${status}\n",
	}))
	e.Use(middleware.Recover())
	//CORS
	e.Use(middleware.CORS())

	// Routes
	e.POST(fmt.Sprintf("/api/%s/auth/register", os.Getenv("api_version")), auth.RegisterHandler)
	e.POST(fmt.Sprintf("/api/%s/auth/login", os.Getenv("api_version")), auth.LoginHandler)
	// Logger
	e.Logger.Fatal(e.Start(":8001"))
}
