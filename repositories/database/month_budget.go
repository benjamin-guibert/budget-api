package database

import (
	"errors"
	"log"

	"github.com/benjamin-guibert/budget-api/models"
	"gorm.io/gorm"
)

type MonthBudgetRepository interface {
	DatabaseRepository
	GetMonthBudget(year, month int) (budget *models.MonthBudget, err error)
}

type MonthBudgetRepositoryHandler struct {
	MonthBudgetRepository
	Database Database
}

func NewMonthBudgetRepository(database Database) *MonthBudgetRepositoryHandler {
	return &MonthBudgetRepositoryHandler{
		Database: database,
	}
}

func (repository *MonthBudgetRepositoryHandler) Seed() error {
	log.Println("Seeding month budgets...")
	budgets := []models.MonthBudget{}
	err := repository.Database.CheckInit()
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	for i := 1; i <= 12; i++ {
		budgets = append(budgets, models.MonthBudget{
			Year:  2021,
			Month: i,
		})
	}

	result := repository.Database.GetGormDB().Create(&budgets)

	log.Println("Month budgets seeded")

	return result.Error
}

func (repository *MonthBudgetRepositoryHandler) GetMonthBudget(
	year, month int) (budget *models.MonthBudget, err error) {
	result := repository.Database.GetGormDB().Preload("BudgetRecords").Where(
		&models.MonthBudget{Year: year, Month: month}).Take(&budget)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return budget, nil
}
