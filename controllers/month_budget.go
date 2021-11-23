package controllers

import (
	"log"
	"net/http"

	"github.com/benjamin-guibert/budget-api/services"
	"github.com/gorilla/mux"
)

type MonthBudgetController interface {
	GetMonthBudget(writer http.ResponseWriter, request *http.Request)
}

type MonthBudgetControllerHandler struct {
	MonthBudgetController
	Service services.MonthBudgetService
}

func NewMonthBudgetController(service services.MonthBudgetService) *MonthBudgetControllerHandler {
	return &MonthBudgetControllerHandler{
		Service: service,
	}
}

func (router MonthBudgetControllerHandler) GetMonthBudget(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	year, err := GetVarToInt(writer, vars, "year")
	if err != nil {
		return
	}
	month, err := GetVarToInt(writer, vars, "month")
	if err != nil {
		return
	}

	budget, err := router.Service.GetMonthBudget(year, month)
	if err != nil {
		log.Println("ERROR:", err)
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}
	if budget == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	ReturnJson(writer, http.StatusOK, budget)
}
