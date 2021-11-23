package models

import (
	"testing"

	"github.com/benjamin-guibert/budget-api/tests"
)

func TestBudgetRecordValidationValid(t *testing.T) {
	record := createBudgetRecord()

	tests.ValidateModel(t, true, record)
}

func TestBudgetRecordValidationMonthBudgetInvalid(t *testing.T) {
	record := createBudgetRecord()
	record.MonthBudget = &MonthBudget{}

	tests.ValidateModel(t, false, record)
}

func TestBudgetRecordValidationMonthLabelInvalid(t *testing.T) {
	record := createBudgetRecord()
	record.Label = ""

	tests.ValidateModel(t, false, record)
}

func TestBudgetRecordValidationMonthAmountInvalid(t *testing.T) {
	record := createBudgetRecord()
	record.Amount = -1

	tests.ValidateModel(t, false, record)
}

func createBudgetRecord() *BudgetRecord {
	return &BudgetRecord{
		MonthBudget: &MonthBudget{
			Year:  2021,
			Month: 6,
		},
		Label:      "Expense #1",
		RecordType: Expense,
		Category:   Needs,
		Amount:     1000,
	}
}
