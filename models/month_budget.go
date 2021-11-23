package models

type MonthBudget struct {
	Model
	Year          int            `gorm:"uniqueIndex:idx_date;<-:create;not null" validate:"required,gte=1000,lte=9999" json:"year"`
	Month         int            `gorm:"uniqueIndex:idx_date;<-:create;not null" validate:"required,gte=1,lte=12" json:"month"`
	BudgetRecords []BudgetRecord `gorm:"constraint:OnDelete:CASCADE;" json:"records"`
}
