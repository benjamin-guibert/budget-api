package services

import (
	"log"

	"github.com/benjamin-guibert/budget-api/models"
	"github.com/benjamin-guibert/budget-api/repositories/database"
)

type MonthBudgetService interface {
	GetMonthBudget(year, month int) (budget *models.MonthBudget, err error)
}

type MonthBudgetServiceHandler struct {
	MonthBudgetService
	MonthBudgetRepository database.MonthBudgetRepository
}

func NewMonthBudgetService(monthBudgetRepository database.MonthBudgetRepository) *MonthBudgetServiceHandler {
	return &MonthBudgetServiceHandler{
		MonthBudgetRepository: monthBudgetRepository,
	}
}

func (service *MonthBudgetServiceHandler) GetMonthBudget(
	year, month int) (budget *models.MonthBudget, err error) {
	budget, err = service.MonthBudgetRepository.GetMonthBudget(year, month)
	if err == nil {
		log.Printf("Month budget loaded: %+v", budget)
	} else {
		log.Println("ERROR:", err)
	}

	return budget, err
}
