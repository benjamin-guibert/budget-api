package database

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gotest.tools/assert"
)

type Repository struct {
	DatabaseRepository
}

type TestModel struct {
	gorm.Model
}

func TestNewDatabase(t *testing.T) {
	database := NewDatabase()

	assert.Assert(t, database != nil)
}

func TestStart(t *testing.T) {
	database := NewDatabase()
	err := database.Start(nil)
	defer StopDatabase(database)

	assert.Assert(t, database.GormDB != nil)
	assert.Assert(t, err == nil, err)
	sqlDB, _ := database.GormDB.DB()
	assert.Assert(t, sqlDB.Ping())
}

func TestStartFailed(t *testing.T) {
	database := NewDatabase()
	password := os.Getenv("DB_PASSWORD")
	os.Setenv("DB_PASSWORD", "xxx")
	defer StopDatabase(database)

	err := database.Start(&gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	assert.Assert(t, err != nil)
	os.Setenv("DB_PASSWORD", password)
}

func TestStop(t *testing.T) {
	database := NewDatabase()
	database.Start(nil)
	defer StopDatabase(database)

	assert.NilError(t, database.Stop())

	assert.Assert(t, database.GormDB == nil)
}

func TestStopAlready(t *testing.T) {
	database := NewDatabase()
	defer StopDatabase(database)

	assert.NilError(t, database.Stop())

	assert.Assert(t, database.GormDB == nil)
}

func TestMigrate(t *testing.T) {
	database := NewDatabase()
	database.Start(nil)
	defer StopDatabase(database)

	assert.NilError(t, database.Migrate(&TestModel{}))
}

func TestMigrateFailed(t *testing.T) {
	database := NewDatabase()
	defer StopDatabase(database)

	err := database.Migrate(&TestModel{})

	assert.Assert(t, err != nil)
}

func StopDatabase(database *DatabaseHandler) {
	database.Stop()
}

func TestGetGormDB(t *testing.T) {
	database := NewDatabase()
	database.Start(&gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	Migrate(database)

	result := database.GetGormDB()

	assert.Assert(t, result != nil)
}

func init() {
	log.SetOutput(ioutil.Discard)
	godotenv.Load(".env.test")
}
