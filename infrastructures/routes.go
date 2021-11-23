package infrastructures

import (
	"net/http"

	"github.com/benjamin-guibert/budget-api/controllers"
)

func AddRoutes(router Router, monthBudgetController controllers.MonthBudgetController) {
	router.AddRoute("/month-budgets/{year}/{month}", monthBudgetController.GetMonthBudget, http.MethodGet)
}
