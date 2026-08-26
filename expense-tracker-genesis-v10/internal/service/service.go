package service

import (
	"fmt"
	"time"

	"expense-tracker/internal/store"
)

type Service struct {
	store *store.ExpenseStore
}

func NewService(s *store.ExpenseStore) *Service {
	return &Service{store: s}
}

type Expense = store.Expense

func (s *Service) AddExpense(amount float64, desc, cat string, date time.Time) (Expense, error) {
	if amount <= 0 {
		return Expense{}, fmt.Errorf("amount must be positive")
	}
	if date.IsZero() {
		date = time.Now()
	}
	return s.store.Add(store.Expense{
		Amount:      amount,
		Description: desc,
		Category:    cat,
		Date:        date,
	})
}

func (s *Service) ListExpenses() ([]Expense, error) {
	return s.store.List()
}

func (s *Service) GetSummary() (map[string]float64, float64, error) {
	return s.store.Summary()
}
