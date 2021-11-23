package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Start(config *gorm.Config) error
	Stop() error
	CheckInit() error
	GetGormDB() *gorm.DB
	Migrate(models ...interface{}) error
}

type DatabaseRepository interface {
	Seed() error
}

type DatabaseHandler struct {
	GormDB *gorm.DB
}

func NewDatabase() *DatabaseHandler {
	return &DatabaseHandler{}
}

func (database *DatabaseHandler) Start(config *gorm.Config) error {
	if config == nil {
		config = &gorm.Config{}
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	log.Println("Starting database connection:", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	database.GormDB = db

	log.Println("Database connection started")

	return nil
}

func (database *DatabaseHandler) Stop() error {
	if database.GormDB == nil {
		return nil
	}
	sqlDB, err := database.GormDB.DB()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}
	database.GormDB = nil

	log.Println("Database connection stopped")

	return sqlDB.Close()
}

func (database *DatabaseHandler) Migrate(models ...interface{}) error {
	err := database.CheckInit()
	if err != nil {
		log.Println("ERROR: ", err)
		return err
	}

	err = database.GormDB.AutoMigrate(models...)
	if err == nil {
		log.Println("Database migrated")
	} else {
		log.Println("ERROR: ", err)
	}

	return err
}

func (database *DatabaseHandler) CheckInit() error {
	if database.GormDB == nil {
		return errors.New("database connection is not initialized")
	}

	return nil
}

func (database *DatabaseHandler) GetGormDB() *gorm.DB {
	return database.GormDB
}
