package database

import (
	"github.com/benjamin-guibert/budget-api/models"
)

func Migrate(database Database) error {
	return database.Migrate(&models.MonthBudget{}, &models.BudgetRecord{})
}
