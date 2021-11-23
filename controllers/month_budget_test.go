package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/benjamin-guibert/budget-api/models"
	"github.com/benjamin-guibert/budget-api/services"
	"github.com/gorilla/mux"
	"gotest.tools/assert"
)

type MonthBudgetServiceMock struct {
	services.MonthBudgetService
}

var MonthBudgetService services.MonthBudgetService
var GetMonthBudgetMock func(year, month int) (budget *models.MonthBudget, err error)
var Controller MonthBudgetController
var Router *mux.Router

func (*MonthBudgetServiceMock) GetMonthBudget(year, month int) (budget *models.MonthBudget, err error) {
	return GetMonthBudgetMock(year, month)
}

func TestNewMonthBudgetController(t *testing.T) {
	controller := NewMonthBudgetController(MonthBudgetService)

	assert.Assert(t, controller != nil)
	assert.Assert(t, controller.Service == MonthBudgetService)
}

func TestGetMonthBudget(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/month-budgets/2021/6", nil)
	recorder := httptest.NewRecorder()
	budget := &models.MonthBudget{Year: 2021, Month: 6}
	GetMonthBudgetMock = func(year, month int) (b *models.MonthBudget, err error) {
		return budget, nil
	}

	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header()["Content-Type"][0])
	decoder := json.NewDecoder(recorder.Body)
	var result models.MonthBudget
	decoder.Decode(&result)
	assert.DeepEqual(t, budget, &result)
}

func TestGetMonthBudgetUnknown(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/month-budgets/2021/6", nil)
	recorder := httptest.NewRecorder()
	GetMonthBudgetMock = func(year, month int) (b *models.MonthBudget, err error) {
		return nil, nil
	}

	Router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Equal(t, 0, recorder.Body.Len())
}

func init() {
	log.SetOutput(ioutil.Discard)
	MonthBudgetService = &MonthBudgetServiceMock{}
	Controller = NewMonthBudgetController(MonthBudgetService)
	Router = mux.NewRouter()
	Router.HandleFunc("/month-budgets/{year}/{month}", Controller.GetMonthBudget).Methods(http.MethodGet)
}
