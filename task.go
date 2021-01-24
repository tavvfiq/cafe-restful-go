package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/tavvfiq/cafe-rest-api-gorm/database"
	"github.com/tavvfiq/cafe-rest-api-gorm/database/model"
	"github.com/tavvfiq/cafe-rest-api-gorm/database/seeder"

	"github.com/tavvfiq/cafe-rest-api-gorm/router"

	"log"
)

// Start start router
func Start() {
	database.SetDatabase(os.Getenv("DB_NAME"))
	router.Start()
}

// Create database
func Create() error {
	dbName := os.Getenv("DB_NAME")
	log.Printf("creating database: %s", dbName)
	result := database.Db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if result.Error != nil {
		return errors.Wrap(result.Error, "error on creating database")
	}
	log.Println("create database success")
	return nil
}

// Delete database
func Delete() error {
	dbName := os.Getenv("DB_NAME")
	log.Printf("deleting database %s", dbName)
	result := database.Db.Exec("DROP DATABASE IF EXISTS " + dbName)
	if result.Error != nil {
		return errors.Wrap(result.Error, "error on deleting database")
	}
	log.Println("delete database success")
	return nil
}

// Migrate database
func Migrate() error {
	dbName := os.Getenv("DB_NAME")
	database.SetDatabase(dbName)
	log.Printf("migrating...")
	error := database.Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Level{}, &model.User{}, &model.Category{}, &model.Menu{}, &model.History{}, &model.OrderHistory{})
	if error != nil {
		return errors.Wrap(error, "error migrating database")
	}
	log.Printf("migrating success")
	return nil
}

// Seed database
func Seed(tableName string) error {
	switch tableName {
	case "user":
		log.Printf("seeding demo-user")
		database.Db.Create(seeder.CreateUser())
	case "level":
		log.Printf("seeding level-level")
		database.Db.Create(seeder.CreateLevel())
	case "category":
		log.Printf("seeding level-category")
		database.Db.Create(seeder.CreateCategory())
	default:
		log.Printf("seeding all")
		database.Db.Create(seeder.CreateLevel())
		database.Db.Create(seeder.CreateUser())
		database.Db.Create(seeder.CreateCategory())
	}
	log.Printf("seeding success")
	return nil
}

// Reset database (drop, create, migrate, seed)
func Reset() {
	// deleting previous database
	if error := Delete(); error != nil {
		log.Println(error)
	}
	// create new database
	if error := Create(); error != nil {
		log.Println(error)
	}
	// migrate tables
	if error := Migrate(); error != nil {
		log.Println(error)
	}
	// seed data
	if error := Seed("all"); error != nil {
		log.Println(error)
	}
}
