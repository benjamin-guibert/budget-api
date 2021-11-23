package models

import (
	"testing"

	"github.com/benjamin-guibert/budget-api/tests"
)

func TestMonthBudgetValidationValid(t *testing.T) {
	budget := createMonthBudget()

	tests.ValidateModel(t, true, budget)
}

func TestMonthBudgetValidationYearInvalid(t *testing.T) {
	budget := createMonthBudget()
	values := []int{999, 10000}

	for _, value := range values {
		budget.Year = value
		tests.ValidateModel(t, false, budget)
	}
}

func TestMonthBudgetValidationMonthInvalid(t *testing.T) {
	budget := createMonthBudget()
	values := []int{0, 13}

	for _, value := range values {
		budget.Month = value
		tests.ValidateModel(t, false, budget)
	}
}

func createMonthBudget() *MonthBudget {
	return &MonthBudget{
		Year:  2021,
		Month: 1,
	}
}
