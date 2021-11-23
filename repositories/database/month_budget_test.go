package database

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/benjamin-guibert/budget-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gotest.tools/assert"
)

var database *DatabaseHandler

func TestNewMonthBudgetRepository(t *testing.T) {
	database := NewDatabase()

	repository := NewMonthBudgetRepository(database)

	assert.Assert(t, repository != nil)
	assert.Assert(t, repository.Database == database)
}

func TestSeed(t *testing.T) {
	prepareDatabase()
	repository := NewMonthBudgetRepository(database)
	var budgetsCount int64
	defer StopDatabase(database)
	defer resetBudgets(database)

	assert.NilError(t, repository.Seed())

	database.GormDB.Model(&models.MonthBudget{}).Count(&budgetsCount)
	assert.Equal(t, int64(12), budgetsCount)
}

func TestSeedAlready(t *testing.T) {
	prepareDatabase()
	repository := NewMonthBudgetRepository(database)
	repository.Seed()
	defer StopDatabase(database)
	defer resetBudgets(database)

	err := repository.Seed()

	assert.Assert(t, err != nil)
}

func TestSeedFailed(t *testing.T) {
	prepareDatabase()
	repository := NewMonthBudgetRepository(database)
	StopDatabase(database)

	err := repository.Seed()

	assert.Assert(t, err != nil)
}

func TestGetMonthBudget(t *testing.T) {
	prepareDatabase()
	repository := NewMonthBudgetRepository(database)
	defer StopDatabase(database)
	defer resetBudgets(database)
	budget := &models.MonthBudget{Year: 2021, Month: 6}
	database.GormDB.Create(budget)

	result, err := repository.GetMonthBudget(2021, 6)

	assert.Assert(t, err == nil, err)
	assert.Equal(t, budget.ID, result.ID)
}

func TestGetMonthBudgetNotFound(t *testing.T) {
	prepareDatabase()
	repository := NewMonthBudgetRepository(database)
	defer StopDatabase(database)
	defer resetBudgets(database)

	result, err := repository.GetMonthBudget(2021, 6)

	assert.Assert(t, err == nil)
	assert.Assert(t, result == nil)
}

func init() {
	log.SetOutput(ioutil.Discard)
	database = NewDatabase()
}

func prepareDatabase() {
	database.Start(&gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	Migrate(database)
	resetBudgets(database)
}

func resetBudgets(database *DatabaseHandler) {
	database.GormDB.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(&models.MonthBudget{})
}
