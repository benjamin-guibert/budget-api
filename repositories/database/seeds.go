package database

func Seed(monthBudgetRepository MonthBudgetRepository) error {
	return monthBudgetRepository.Seed()
}
