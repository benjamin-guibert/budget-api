package models

type BudgetRecordType int

const (
	Expense BudgetRecordType = 0
	Income  BudgetRecordType = 1
)

func (recordType BudgetRecordType) String() string {
	switch recordType {
	case Expense:
		return "expense"
	case Income:
		return "income"
	}
	return "unknown"
}

type BudgetRecordCategory int

const (
	Needs BudgetRecordCategory = 0
	Wants BudgetRecordCategory = 1
)

func (category BudgetRecordCategory) String() string {
	switch category {
	case Needs:
		return "needs"
	case Wants:
		return "wants"
	}
	return "unknown"
}

type BudgetRecord struct {
	Model
	MonthBudgetID uint                 `gorm:"not null" json:"-"`
	MonthBudget   *MonthBudget         `validate:"required" json:"budget,omitempty"`
	Label         string               `gorm:"not null" validate:"required" json:"label"`
	RecordType    BudgetRecordType     `gorm:"not null" json:"type"`
	Category      BudgetRecordCategory `gorm:"not null" json:"category"`
	Amount        int                  `gorm:"not null" validate:"required,gte=0" json:"amount"`
}
