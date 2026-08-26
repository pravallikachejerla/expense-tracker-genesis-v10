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

func (s *Service) GetExpense(id int) (Expense, error) {
	return s.store.GetByID(id)
}

func (s *Service) ListExpenses() ([]Expense, error) {
	return s.store.List()
}

func (s *Service) ListByCategory(cat string) ([]Expense, error) {
	return s.store.ListByCategory(cat)
}

func (s *Service) ListByDateRange(start, end time.Time) ([]Expense, error) {
	return s.store.ListByDateRange(start, end)
}

func (s *Service) UpdateExpense(id int, amount float64, desc, cat string, date time.Time) (Expense, error) {
	if amount <= 0 {
		return Expense{}, fmt.Errorf("amount must be positive")
	}
	if date.IsZero() {
		date = time.Now()
	}
	return s.store.Update(store.Expense{
		ID:          id,
		Amount:      amount,
		Description: desc,
		Category:    cat,
		Date:        date,
	})
}

func (s *Service) DeleteExpense(id int) error {
	return s.store.Delete(id)
}

func (s *Service) GetSummary() (map[string]float64, float64, error) {
	return s.store.Summary()
}

func (s *Service) GetCategories() ([]string, error) {
	return s.store.GetCategories()
}

func (s *Service) AddCategory(name string) error {
	if name == "" {
		return fmt.Errorf("category name cannot be empty")
	}
	return s.store.AddCategory(name)
}
