package services

import (
	"errors"
	"io/ioutil"
	"log"
	"testing"

	"github.com/benjamin-guibert/budget-api/models"
	"github.com/benjamin-guibert/budget-api/repositories/database"
	"gotest.tools/assert"
)

type MonthBudgetRepositoryMock struct {
	database.MonthBudgetRepository
}

var MonthBudgetRepository database.MonthBudgetRepository
var GetMonthBudgetMock func(year, month int) (budget *models.MonthBudget, err error)

func (*MonthBudgetRepositoryMock) GetMonthBudget(year, month int) (budget *models.MonthBudget, err error) {
	return GetMonthBudgetMock(year, month)
}

func TestNewMonthBudgetService(t *testing.T) {
	service := NewMonthBudgetService(MonthBudgetRepository)

	assert.Assert(t, service != nil)
	assert.Assert(t, service.MonthBudgetRepository == MonthBudgetRepository)
}

func TestGetMonthBudget(t *testing.T) {
	service := NewMonthBudgetService(MonthBudgetRepository)
	budget := &models.MonthBudget{Year: 2021, Month: 6}
	GetMonthBudgetMock = func(year, month int) (b *models.MonthBudget, err error) {
		return budget, nil
	}

	result, err := service.GetMonthBudget(2021, 6)

	assert.Assert(t, result == budget)
	assert.Assert(t, err == nil, err)
}

func TestGetMonthBudgetFailed(t *testing.T) {
	service := NewMonthBudgetService(MonthBudgetRepository)
	GetMonthBudgetMock = func(year, month int) (b *models.MonthBudget, err error) {
		return nil, errors.New("error")
	}

	result, err := service.GetMonthBudget(2021, 6)

	assert.Assert(t, result == nil)
	assert.Assert(t, err != nil)
}

func init() {
	log.SetOutput(ioutil.Discard)
	MonthBudgetRepository = &MonthBudgetRepositoryMock{}
}
