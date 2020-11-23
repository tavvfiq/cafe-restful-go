package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

// Db exported Db variable
var Db *gorm.DB

func init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// get variable from .env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")

	// connect to database
	Db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/?parseTime=true", dbUser, dbPass, dbHost)), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	sqlDB, err := Db.DB()

	if err != nil {
		log.Fatal(err)
	} else {
		if err := sqlDB.PingContext(ctx); err != nil {
			log.Fatal(err)
		} else {
			log.Println("Connected to mysql database")
		}
	}
}

// SetDatabase set database name
func SetDatabase(dbName string) {
	Db.Exec("USE " + dbName)
}
